// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the information about the registered snaphosts repositories
func LookupSnapshotRepository(ctx *pulumi.Context, args *LookupSnapshotRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotRepositoryResult, error) {
	var rv LookupSnapshotRepositoryResult
	err := ctx.Invoke("elasticstack:index/getSnapshotRepository:getSnapshotRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshotRepository.
type LookupSnapshotRepositoryArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *GetSnapshotRepositoryElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Name of the snapshot repository.
	Name string `pulumi:"name"`
}

// A collection of values returned by getSnapshotRepository.
type LookupSnapshotRepositoryResult struct {
	// Azure Blob storage as a repository. Set only if the type of the fetched repo is `azure`.
	Azures []GetSnapshotRepositoryAzure `pulumi:"azures"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *GetSnapshotRepositoryElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Shared filesystem repository. Set only if the type of the fetched repo is `fs`.
	Fs []GetSnapshotRepositoryF `pulumi:"fs"`
	// Google Cloud Storage service as a repository. Set only if the type of the fetched repo is `gcs`.
	Gcs []GetSnapshotRepositoryGc `pulumi:"gcs"`
	// HDFS File System as a repository. Set only if the type of the fetched repo is `hdfs`.
	Hdfs []GetSnapshotRepositoryHdf `pulumi:"hdfs"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Name of the snapshot repository.
	Name string `pulumi:"name"`
	// AWS S3 as a repository. Set only if the type of the fetched repo is `s3`.
	S3s []GetSnapshotRepositoryS3 `pulumi:"s3s"`
	// Repository type.
	Type string `pulumi:"type"`
	// URL repository. Set only if the type of the fetched repo is `url`.
	Urls []GetSnapshotRepositoryUrl `pulumi:"urls"`
}

func LookupSnapshotRepositoryOutput(ctx *pulumi.Context, args LookupSnapshotRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotRepositoryResult, error) {
			args := v.(LookupSnapshotRepositoryArgs)
			r, err := LookupSnapshotRepository(ctx, &args, opts...)
			var s LookupSnapshotRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotRepositoryResultOutput)
}

// A collection of arguments for invoking getSnapshotRepository.
type LookupSnapshotRepositoryOutputArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection GetSnapshotRepositoryElasticsearchConnectionPtrInput `pulumi:"elasticsearchConnection"`
	// Name of the snapshot repository.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupSnapshotRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getSnapshotRepository.
type LookupSnapshotRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotRepositoryResult)(nil)).Elem()
}

func (o LookupSnapshotRepositoryResultOutput) ToLookupSnapshotRepositoryResultOutput() LookupSnapshotRepositoryResultOutput {
	return o
}

func (o LookupSnapshotRepositoryResultOutput) ToLookupSnapshotRepositoryResultOutputWithContext(ctx context.Context) LookupSnapshotRepositoryResultOutput {
	return o
}

// Azure Blob storage as a repository. Set only if the type of the fetched repo is `azure`.
func (o LookupSnapshotRepositoryResultOutput) Azures() GetSnapshotRepositoryAzureArrayOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) []GetSnapshotRepositoryAzure { return v.Azures }).(GetSnapshotRepositoryAzureArrayOutput)
}

// Elasticsearch connection configuration block.
func (o LookupSnapshotRepositoryResultOutput) ElasticsearchConnection() GetSnapshotRepositoryElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) *GetSnapshotRepositoryElasticsearchConnection {
		return v.ElasticsearchConnection
	}).(GetSnapshotRepositoryElasticsearchConnectionPtrOutput)
}

// Shared filesystem repository. Set only if the type of the fetched repo is `fs`.
func (o LookupSnapshotRepositoryResultOutput) Fs() GetSnapshotRepositoryFArrayOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) []GetSnapshotRepositoryF { return v.Fs }).(GetSnapshotRepositoryFArrayOutput)
}

// Google Cloud Storage service as a repository. Set only if the type of the fetched repo is `gcs`.
func (o LookupSnapshotRepositoryResultOutput) Gcs() GetSnapshotRepositoryGcArrayOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) []GetSnapshotRepositoryGc { return v.Gcs }).(GetSnapshotRepositoryGcArrayOutput)
}

// HDFS File System as a repository. Set only if the type of the fetched repo is `hdfs`.
func (o LookupSnapshotRepositoryResultOutput) Hdfs() GetSnapshotRepositoryHdfArrayOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) []GetSnapshotRepositoryHdf { return v.Hdfs }).(GetSnapshotRepositoryHdfArrayOutput)
}

// Internal identifier of the resource
func (o LookupSnapshotRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the snapshot repository.
func (o LookupSnapshotRepositoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) string { return v.Name }).(pulumi.StringOutput)
}

// AWS S3 as a repository. Set only if the type of the fetched repo is `s3`.
func (o LookupSnapshotRepositoryResultOutput) S3s() GetSnapshotRepositoryS3ArrayOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) []GetSnapshotRepositoryS3 { return v.S3s }).(GetSnapshotRepositoryS3ArrayOutput)
}

// Repository type.
func (o LookupSnapshotRepositoryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) string { return v.Type }).(pulumi.StringOutput)
}

// URL repository. Set only if the type of the fetched repo is `url`.
func (o LookupSnapshotRepositoryResultOutput) Urls() GetSnapshotRepositoryUrlArrayOutput {
	return o.ApplyT(func(v LookupSnapshotRepositoryResult) []GetSnapshotRepositoryUrl { return v.Urls }).(GetSnapshotRepositoryUrlArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotRepositoryResultOutput{})
}