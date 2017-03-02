package kubernetes

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	api "k8s.io/kubernetes/pkg/api/v1"
)

func resourceKubernetesConfigMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesConfigMapCreate,
		Read:   resourceKubernetesConfigMapRead,
		Update: resourceKubernetesConfigMapUpdate,
		Delete: resourceKubernetesConfigMapDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"metadata": metadataSchema,
			"data": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceKubernetesConfigMapCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	configMap := api.Confi{
		ObjectMeta: metadata,
	}
	log.Printf("[INFO] Creating new namespace: %#v", namespace)
	out, err := conn.CoreV1().Namespaces().Create(&namespace)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Submitted new namespace: %#v", out)
	d.SetId(out.Name)

	return resourceKubernetesConfigMapRead(d, meta)
}

func resourceKubernetesConfigMapRead(d *schema.ResourceData, meta interface{}) error {

}

func resourceKubernetesConfigMapUpdate(d *schema.ResourceData, meta interface{}) error {

}

func resourceKubernetesConfigMapDelete(d *schema.ResourceData, meta interface{}) error {

}
