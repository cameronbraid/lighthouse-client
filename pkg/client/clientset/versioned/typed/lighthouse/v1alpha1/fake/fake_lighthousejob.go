// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/jenkins-x/lighthouse-client/pkg/apis/lighthouse/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLighthouseJobs implements LighthouseJobInterface
type FakeLighthouseJobs struct {
	Fake *FakeLighthouseV1alpha1
	ns   string
}

var lighthousejobsResource = schema.GroupVersionResource{Group: "lighthouse.jenkins.io", Version: "v1alpha1", Resource: "lighthousejobs"}

var lighthousejobsKind = schema.GroupVersionKind{Group: "lighthouse.jenkins.io", Version: "v1alpha1", Kind: "LighthouseJob"}

// Get takes name of the lighthouseJob, and returns the corresponding lighthouseJob object, and an error if there is any.
func (c *FakeLighthouseJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LighthouseJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lighthousejobsResource, c.ns, name), &v1alpha1.LighthouseJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LighthouseJob), err
}

// List takes label and field selectors, and returns the list of LighthouseJobs that match those selectors.
func (c *FakeLighthouseJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LighthouseJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lighthousejobsResource, lighthousejobsKind, c.ns, opts), &v1alpha1.LighthouseJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LighthouseJobList{ListMeta: obj.(*v1alpha1.LighthouseJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.LighthouseJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lighthouseJobs.
func (c *FakeLighthouseJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lighthousejobsResource, c.ns, opts))

}

// Create takes the representation of a lighthouseJob and creates it.  Returns the server's representation of the lighthouseJob, and an error, if there is any.
func (c *FakeLighthouseJobs) Create(ctx context.Context, lighthouseJob *v1alpha1.LighthouseJob, opts v1.CreateOptions) (result *v1alpha1.LighthouseJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lighthousejobsResource, c.ns, lighthouseJob), &v1alpha1.LighthouseJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LighthouseJob), err
}

// Update takes the representation of a lighthouseJob and updates it. Returns the server's representation of the lighthouseJob, and an error, if there is any.
func (c *FakeLighthouseJobs) Update(ctx context.Context, lighthouseJob *v1alpha1.LighthouseJob, opts v1.UpdateOptions) (result *v1alpha1.LighthouseJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lighthousejobsResource, c.ns, lighthouseJob), &v1alpha1.LighthouseJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LighthouseJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLighthouseJobs) UpdateStatus(ctx context.Context, lighthouseJob *v1alpha1.LighthouseJob, opts v1.UpdateOptions) (*v1alpha1.LighthouseJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lighthousejobsResource, "status", c.ns, lighthouseJob), &v1alpha1.LighthouseJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LighthouseJob), err
}

// Delete takes name of the lighthouseJob and deletes it. Returns an error if one occurs.
func (c *FakeLighthouseJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lighthousejobsResource, c.ns, name), &v1alpha1.LighthouseJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLighthouseJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lighthousejobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LighthouseJobList{})
	return err
}

// Patch applies the patch and returns the patched lighthouseJob.
func (c *FakeLighthouseJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LighthouseJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lighthousejobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LighthouseJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LighthouseJob), err
}
