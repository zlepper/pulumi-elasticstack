// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates or updates a snapshot lifecycle policy. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-put-policy.html
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/zlepper/pulumi-elasticstack/sdk/go/elasticstack"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		repo, err := elasticstack.NewSnapshotRepository(ctx, "repo", &elasticstack.SnapshotRepositoryArgs{
// 			Fs: &elasticstack.SnapshotRepositoryFsArgs{
// 				Location:              pulumi.String("/tmp/snapshots"),
// 				Compress:              pulumi.Bool(true),
// 				MaxRestoreBytesPerSec: pulumi.String("20mb"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticstack.NewSnapshotLifecycle(ctx, "slmPolicy", &elasticstack.SnapshotLifecycleArgs{
// 			Schedule:     pulumi.String("0 30 1 * * ?"),
// 			SnapshotName: pulumi.String("<daily-snap-{now/d}>"),
// 			Repository:   repo.Name,
// 			Indices: pulumi.StringArray{
// 				pulumi.String("data-*"),
// 				pulumi.String("important"),
// 			},
// 			IgnoreUnavailable:  pulumi.Bool(false),
// 			IncludeGlobalState: pulumi.Bool(false),
// 			ExpireAfter:        pulumi.String("30d"),
// 			MinCount:           pulumi.Int(5),
// 			MaxCount:           pulumi.Int(50),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import elasticstack:index/snapshotLifecycle:SnapshotLifecycle my_policy <cluster_uuid>/<slm policy name>
// ```
type SnapshotLifecycle struct {
	pulumi.CustomResourceState

	// Elasticsearch connection configuration block.
	ElasticsearchConnection SnapshotLifecycleElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards pulumi.StringPtrOutput `pulumi:"expandWildcards"`
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter pulumi.StringPtrOutput `pulumi:"expireAfter"`
	// Feature states to include in the snapshot.
	FeatureStates pulumi.StringArrayOutput `pulumi:"featureStates"`
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable pulumi.BoolPtrOutput `pulumi:"ignoreUnavailable"`
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState pulumi.BoolPtrOutput `pulumi:"includeGlobalState"`
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices pulumi.StringArrayOutput `pulumi:"indices"`
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount pulumi.IntPtrOutput `pulumi:"maxCount"`
	// Attaches arbitrary metadata to the snapshot.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount pulumi.IntPtrOutput `pulumi:"minCount"`
	// ID for the snapshot lifecycle policy you want to create or update.
	Name pulumi.StringOutput `pulumi:"name"`
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial pulumi.BoolPtrOutput `pulumi:"partial"`
	// Repository used to store snapshots created by this policy.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName pulumi.StringPtrOutput `pulumi:"snapshotName"`
}

// NewSnapshotLifecycle registers a new resource with the given unique name, arguments, and options.
func NewSnapshotLifecycle(ctx *pulumi.Context,
	name string, args *SnapshotLifecycleArgs, opts ...pulumi.ResourceOption) (*SnapshotLifecycle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	var resource SnapshotLifecycle
	err := ctx.RegisterResource("elasticstack:index/snapshotLifecycle:SnapshotLifecycle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotLifecycle gets an existing SnapshotLifecycle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotLifecycle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotLifecycleState, opts ...pulumi.ResourceOption) (*SnapshotLifecycle, error) {
	var resource SnapshotLifecycle
	err := ctx.ReadResource("elasticstack:index/snapshotLifecycle:SnapshotLifecycle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotLifecycle resources.
type snapshotLifecycleState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *SnapshotLifecycleElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards *string `pulumi:"expandWildcards"`
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter *string `pulumi:"expireAfter"`
	// Feature states to include in the snapshot.
	FeatureStates []string `pulumi:"featureStates"`
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable *bool `pulumi:"ignoreUnavailable"`
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState *bool `pulumi:"includeGlobalState"`
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices []string `pulumi:"indices"`
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount *int `pulumi:"maxCount"`
	// Attaches arbitrary metadata to the snapshot.
	Metadata *string `pulumi:"metadata"`
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount *int `pulumi:"minCount"`
	// ID for the snapshot lifecycle policy you want to create or update.
	Name *string `pulumi:"name"`
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial *bool `pulumi:"partial"`
	// Repository used to store snapshots created by this policy.
	Repository *string `pulumi:"repository"`
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule *string `pulumi:"schedule"`
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName *string `pulumi:"snapshotName"`
}

type SnapshotLifecycleState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection SnapshotLifecycleElasticsearchConnectionPtrInput
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards pulumi.StringPtrInput
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter pulumi.StringPtrInput
	// Feature states to include in the snapshot.
	FeatureStates pulumi.StringArrayInput
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable pulumi.BoolPtrInput
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState pulumi.BoolPtrInput
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices pulumi.StringArrayInput
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount pulumi.IntPtrInput
	// Attaches arbitrary metadata to the snapshot.
	Metadata pulumi.StringPtrInput
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount pulumi.IntPtrInput
	// ID for the snapshot lifecycle policy you want to create or update.
	Name pulumi.StringPtrInput
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial pulumi.BoolPtrInput
	// Repository used to store snapshots created by this policy.
	Repository pulumi.StringPtrInput
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule pulumi.StringPtrInput
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName pulumi.StringPtrInput
}

func (SnapshotLifecycleState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotLifecycleState)(nil)).Elem()
}

type snapshotLifecycleArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *SnapshotLifecycleElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards *string `pulumi:"expandWildcards"`
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter *string `pulumi:"expireAfter"`
	// Feature states to include in the snapshot.
	FeatureStates []string `pulumi:"featureStates"`
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable *bool `pulumi:"ignoreUnavailable"`
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState *bool `pulumi:"includeGlobalState"`
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices []string `pulumi:"indices"`
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount *int `pulumi:"maxCount"`
	// Attaches arbitrary metadata to the snapshot.
	Metadata *string `pulumi:"metadata"`
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount *int `pulumi:"minCount"`
	// ID for the snapshot lifecycle policy you want to create or update.
	Name *string `pulumi:"name"`
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial *bool `pulumi:"partial"`
	// Repository used to store snapshots created by this policy.
	Repository string `pulumi:"repository"`
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule string `pulumi:"schedule"`
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName *string `pulumi:"snapshotName"`
}

// The set of arguments for constructing a SnapshotLifecycle resource.
type SnapshotLifecycleArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection SnapshotLifecycleElasticsearchConnectionPtrInput
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards pulumi.StringPtrInput
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter pulumi.StringPtrInput
	// Feature states to include in the snapshot.
	FeatureStates pulumi.StringArrayInput
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable pulumi.BoolPtrInput
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState pulumi.BoolPtrInput
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices pulumi.StringArrayInput
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount pulumi.IntPtrInput
	// Attaches arbitrary metadata to the snapshot.
	Metadata pulumi.StringPtrInput
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount pulumi.IntPtrInput
	// ID for the snapshot lifecycle policy you want to create or update.
	Name pulumi.StringPtrInput
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial pulumi.BoolPtrInput
	// Repository used to store snapshots created by this policy.
	Repository pulumi.StringInput
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule pulumi.StringInput
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName pulumi.StringPtrInput
}

func (SnapshotLifecycleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotLifecycleArgs)(nil)).Elem()
}

type SnapshotLifecycleInput interface {
	pulumi.Input

	ToSnapshotLifecycleOutput() SnapshotLifecycleOutput
	ToSnapshotLifecycleOutputWithContext(ctx context.Context) SnapshotLifecycleOutput
}

func (*SnapshotLifecycle) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotLifecycle)(nil)).Elem()
}

func (i *SnapshotLifecycle) ToSnapshotLifecycleOutput() SnapshotLifecycleOutput {
	return i.ToSnapshotLifecycleOutputWithContext(context.Background())
}

func (i *SnapshotLifecycle) ToSnapshotLifecycleOutputWithContext(ctx context.Context) SnapshotLifecycleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotLifecycleOutput)
}

// SnapshotLifecycleArrayInput is an input type that accepts SnapshotLifecycleArray and SnapshotLifecycleArrayOutput values.
// You can construct a concrete instance of `SnapshotLifecycleArrayInput` via:
//
//          SnapshotLifecycleArray{ SnapshotLifecycleArgs{...} }
type SnapshotLifecycleArrayInput interface {
	pulumi.Input

	ToSnapshotLifecycleArrayOutput() SnapshotLifecycleArrayOutput
	ToSnapshotLifecycleArrayOutputWithContext(context.Context) SnapshotLifecycleArrayOutput
}

type SnapshotLifecycleArray []SnapshotLifecycleInput

func (SnapshotLifecycleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotLifecycle)(nil)).Elem()
}

func (i SnapshotLifecycleArray) ToSnapshotLifecycleArrayOutput() SnapshotLifecycleArrayOutput {
	return i.ToSnapshotLifecycleArrayOutputWithContext(context.Background())
}

func (i SnapshotLifecycleArray) ToSnapshotLifecycleArrayOutputWithContext(ctx context.Context) SnapshotLifecycleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotLifecycleArrayOutput)
}

// SnapshotLifecycleMapInput is an input type that accepts SnapshotLifecycleMap and SnapshotLifecycleMapOutput values.
// You can construct a concrete instance of `SnapshotLifecycleMapInput` via:
//
//          SnapshotLifecycleMap{ "key": SnapshotLifecycleArgs{...} }
type SnapshotLifecycleMapInput interface {
	pulumi.Input

	ToSnapshotLifecycleMapOutput() SnapshotLifecycleMapOutput
	ToSnapshotLifecycleMapOutputWithContext(context.Context) SnapshotLifecycleMapOutput
}

type SnapshotLifecycleMap map[string]SnapshotLifecycleInput

func (SnapshotLifecycleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotLifecycle)(nil)).Elem()
}

func (i SnapshotLifecycleMap) ToSnapshotLifecycleMapOutput() SnapshotLifecycleMapOutput {
	return i.ToSnapshotLifecycleMapOutputWithContext(context.Background())
}

func (i SnapshotLifecycleMap) ToSnapshotLifecycleMapOutputWithContext(ctx context.Context) SnapshotLifecycleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotLifecycleMapOutput)
}

type SnapshotLifecycleOutput struct{ *pulumi.OutputState }

func (SnapshotLifecycleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotLifecycle)(nil)).Elem()
}

func (o SnapshotLifecycleOutput) ToSnapshotLifecycleOutput() SnapshotLifecycleOutput {
	return o
}

func (o SnapshotLifecycleOutput) ToSnapshotLifecycleOutputWithContext(ctx context.Context) SnapshotLifecycleOutput {
	return o
}

// Elasticsearch connection configuration block.
func (o SnapshotLifecycleOutput) ElasticsearchConnection() SnapshotLifecycleElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) SnapshotLifecycleElasticsearchConnectionPtrOutput {
		return v.ElasticsearchConnection
	}).(SnapshotLifecycleElasticsearchConnectionPtrOutput)
}

// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
func (o SnapshotLifecycleOutput) ExpandWildcards() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringPtrOutput { return v.ExpandWildcards }).(pulumi.StringPtrOutput)
}

// Time period after which a snapshot is considered expired and eligible for deletion.
func (o SnapshotLifecycleOutput) ExpireAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringPtrOutput { return v.ExpireAfter }).(pulumi.StringPtrOutput)
}

// Feature states to include in the snapshot.
func (o SnapshotLifecycleOutput) FeatureStates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringArrayOutput { return v.FeatureStates }).(pulumi.StringArrayOutput)
}

// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
func (o SnapshotLifecycleOutput) IgnoreUnavailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.BoolPtrOutput { return v.IgnoreUnavailable }).(pulumi.BoolPtrOutput)
}

// If `true`, include the cluster state in the snapshot.
func (o SnapshotLifecycleOutput) IncludeGlobalState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.BoolPtrOutput { return v.IncludeGlobalState }).(pulumi.BoolPtrOutput)
}

// Comma-separated list of data streams and indices to include in the snapshot.
func (o SnapshotLifecycleOutput) Indices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringArrayOutput { return v.Indices }).(pulumi.StringArrayOutput)
}

// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
func (o SnapshotLifecycleOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.IntPtrOutput { return v.MaxCount }).(pulumi.IntPtrOutput)
}

// Attaches arbitrary metadata to the snapshot.
func (o SnapshotLifecycleOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringOutput { return v.Metadata }).(pulumi.StringOutput)
}

// Minimum number of snapshots to retain, even if the snapshots have expired.
func (o SnapshotLifecycleOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.IntPtrOutput { return v.MinCount }).(pulumi.IntPtrOutput)
}

// ID for the snapshot lifecycle policy you want to create or update.
func (o SnapshotLifecycleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
func (o SnapshotLifecycleOutput) Partial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.BoolPtrOutput { return v.Partial }).(pulumi.BoolPtrOutput)
}

// Repository used to store snapshots created by this policy.
func (o SnapshotLifecycleOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// Periodic or absolute schedule at which the policy creates snapshots.
func (o SnapshotLifecycleOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// Name automatically assigned to each snapshot created by the policy.
func (o SnapshotLifecycleOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotLifecycle) pulumi.StringPtrOutput { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

type SnapshotLifecycleArrayOutput struct{ *pulumi.OutputState }

func (SnapshotLifecycleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotLifecycle)(nil)).Elem()
}

func (o SnapshotLifecycleArrayOutput) ToSnapshotLifecycleArrayOutput() SnapshotLifecycleArrayOutput {
	return o
}

func (o SnapshotLifecycleArrayOutput) ToSnapshotLifecycleArrayOutputWithContext(ctx context.Context) SnapshotLifecycleArrayOutput {
	return o
}

func (o SnapshotLifecycleArrayOutput) Index(i pulumi.IntInput) SnapshotLifecycleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotLifecycle {
		return vs[0].([]*SnapshotLifecycle)[vs[1].(int)]
	}).(SnapshotLifecycleOutput)
}

type SnapshotLifecycleMapOutput struct{ *pulumi.OutputState }

func (SnapshotLifecycleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotLifecycle)(nil)).Elem()
}

func (o SnapshotLifecycleMapOutput) ToSnapshotLifecycleMapOutput() SnapshotLifecycleMapOutput {
	return o
}

func (o SnapshotLifecycleMapOutput) ToSnapshotLifecycleMapOutputWithContext(ctx context.Context) SnapshotLifecycleMapOutput {
	return o
}

func (o SnapshotLifecycleMapOutput) MapIndex(k pulumi.StringInput) SnapshotLifecycleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotLifecycle {
		return vs[0].(map[string]*SnapshotLifecycle)[vs[1].(string)]
	}).(SnapshotLifecycleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotLifecycleInput)(nil)).Elem(), &SnapshotLifecycle{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotLifecycleArrayInput)(nil)).Elem(), SnapshotLifecycleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotLifecycleMapInput)(nil)).Elem(), SnapshotLifecycleMap{})
	pulumi.RegisterOutputType(SnapshotLifecycleOutput{})
	pulumi.RegisterOutputType(SnapshotLifecycleArrayOutput{})
	pulumi.RegisterOutputType(SnapshotLifecycleMapOutput{})
}
