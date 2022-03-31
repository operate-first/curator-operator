/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	curatorv1alpha1 "github.com/operate-first/curator-operator/api/v1alpha1"
)

// FetchDataReconciler reconciles a FetchData object
type FetchDataReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=my.domain,resources=FetchData,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=my.domain,resources=FetchData/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=my.domain,resources=FetchData/finalizers,verbs=update
//+kubebuilder:rbac:groups=batch,resources=cronjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=cronjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=batch,resources=cronjobs/finalizers,verbs=update
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs/status,verbs=get
//+kubebuilder:rbac:groups="core",resources=persistentvolumeclaims,verbs=get;watch;list;create;update;delete
//+kubebuilder:rbac:groups="core",resources=persistentvolumeclaims,verbs=get;watch;list
//+kubebuilder:rbac:groups="core",resources=persistentvolumeclaims/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the FetchData object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile

func (r *FetchDataReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	fd := &curatorv1alpha1.FetchData{}
	if err := r.Get(ctx, req.NamespacedName, fd); err != nil {
		l.Info("failed to query for the FetchData resource", "error", err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// TODO: pivot towards a CreateOrRecreate implementation for the underlying CronJob resource
	if err := r.createCronJob(ctx, fd); err != nil {
		l.Error(err, "failed to create the CronJob resource")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *FetchDataReconciler) createCronJob(ctx context.Context, m *curatorv1alpha1.FetchData) error {
	job := &batchv1.CronJob{}
	err := r.Client.Get(ctx, types.NamespacedName{Name: m.Name, Namespace: m.Namespace}, job)
	if apierrors.IsNotFound(err) {
		job = NewCronJob(m)
		if err := r.Client.Create(ctx, job); err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func NewCronJob(m *curatorv1alpha1.FetchData) *batchv1.CronJob {
	cronjob := &batchv1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind: "Cronjob",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Spec.CronjobNamespace,
		},
		Spec: batchv1.CronJobSpec{
			ConcurrencyPolicy: "Forbid",
			Schedule:          m.Spec.Schedule,
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Volumes: []corev1.Volume{
								{
									Name: "koku-metrics-operator-data",
									VolumeSource: corev1.VolumeSource{
										PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
											ClaimName: "koku-metrics-operator-data",
										},
									},
								},
							},
							Containers: []corev1.Container{
								{
									Name:  "crdunzip",
									Image: "docker.io/surbhi0129/crd_unzip:latest",
									Env: []corev1.EnvVar{
										{
											Name:  "BACKUP_SRC",
											Value: m.Spec.BackupSrc,
										},
										{
											Name:  "UNZIP_DIR",
											Value: m.Spec.UnzipDir,
										},
										{
											Name:  "DATABASE_NAME",
											Value: m.Spec.DatabaseName,
										},
										{
											Name:  "DATABASE_USER",
											Value: m.Spec.DatabaseUser,
										},
										{
											Name:  "DATABASE_PASSWORD",
											Value: m.Spec.DatabasePassword,
										},
										{
											Name:  "DATABASE_HOST_NAME",
											Value: m.Spec.DatabaseHostName,
										},
										{
											Name:  "PORT_NUMBER",
											Value: m.Spec.DatabasePort,
										},
									},
									Command: []string{"python3"},
									Args:    []string{"scripts/unzip_backup.py"},
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      "koku-metrics-operator-data",
											MountPath: "/tmp/koku-metrics-operator-data",
										},
									},
								},
							},
							RestartPolicy: "Never",
						},
					},
				},
			},
		},
	}

	return cronjob
}

// SetupWithManager sets up the controller with the Manager.
func (r *FetchDataReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&curatorv1alpha1.FetchData{}).
		Complete(r)
}
