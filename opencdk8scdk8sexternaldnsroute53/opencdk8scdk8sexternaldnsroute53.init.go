package opencdk8scdk8sexternaldnsroute53

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Affinity",
		reflect.TypeOf((*Affinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AggregationRule",
		reflect.TypeOf((*AggregationRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AggregationRuleV1Alpha1",
		reflect.TypeOf((*AggregationRuleV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AggregationRuleV1Beta1",
		reflect.TypeOf((*AggregationRuleV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AllowedCsiDriverV1Beta1",
		reflect.TypeOf((*AllowedCsiDriverV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AllowedFlexVolumeV1Beta1",
		reflect.TypeOf((*AllowedFlexVolumeV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AllowedHostPathV1Beta1",
		reflect.TypeOf((*AllowedHostPathV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ApiServiceSpec",
		reflect.TypeOf((*ApiServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ApiServiceSpecV1Beta1",
		reflect.TypeOf((*ApiServiceSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AwsElasticBlockStoreVolumeSource",
		reflect.TypeOf((*AwsElasticBlockStoreVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.AwsExternalDns",
		reflect.TypeOf((*AwsExternalDns)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentName", GoGetter: "DeploymentName"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "kongTCP", GoGetter: "KongTCP"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeSelector", GoGetter: "NodeSelector"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountName", GoGetter: "ServiceAccountName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsExternalDns{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AwsExternalDnsOptions",
		reflect.TypeOf((*AwsExternalDnsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.AwsExternalDnsPolicyHelper",
		reflect.TypeOf((*AwsExternalDnsPolicyHelper)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AwsExternalDnsPolicyHelper{}
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AzureDiskVolumeSource",
		reflect.TypeOf((*AzureDiskVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AzureFilePersistentVolumeSource",
		reflect.TypeOf((*AzureFilePersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.AzureFileVolumeSource",
		reflect.TypeOf((*AzureFileVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.BoundObjectReference",
		reflect.TypeOf((*BoundObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Capabilities",
		reflect.TypeOf((*Capabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CephFsPersistentVolumeSource",
		reflect.TypeOf((*CephFsPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CephFsVolumeSource",
		reflect.TypeOf((*CephFsVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CertificateSigningRequestSpec",
		reflect.TypeOf((*CertificateSigningRequestSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CertificateSigningRequestSpecV1Beta1",
		reflect.TypeOf((*CertificateSigningRequestSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CinderPersistentVolumeSource",
		reflect.TypeOf((*CinderPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CinderVolumeSource",
		reflect.TypeOf((*CinderVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ClientIpConfig",
		reflect.TypeOf((*ClientIpConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ComponentCondition",
		reflect.TypeOf((*ComponentCondition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ConfigMapEnvSource",
		reflect.TypeOf((*ConfigMapEnvSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ConfigMapKeySelector",
		reflect.TypeOf((*ConfigMapKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ConfigMapNodeConfigSource",
		reflect.TypeOf((*ConfigMapNodeConfigSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ConfigMapProjection",
		reflect.TypeOf((*ConfigMapProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ConfigMapVolumeSource",
		reflect.TypeOf((*ConfigMapVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Container",
		reflect.TypeOf((*Container)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ContainerPort",
		reflect.TypeOf((*ContainerPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ContainerResourceMetricSourceV2Beta1",
		reflect.TypeOf((*ContainerResourceMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ContainerResourceMetricSourceV2Beta2",
		reflect.TypeOf((*ContainerResourceMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CronJobSpec",
		reflect.TypeOf((*CronJobSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CronJobSpecV1Beta1",
		reflect.TypeOf((*CronJobSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CrossVersionObjectReference",
		reflect.TypeOf((*CrossVersionObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CrossVersionObjectReferenceV2Beta1",
		reflect.TypeOf((*CrossVersionObjectReferenceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CrossVersionObjectReferenceV2Beta2",
		reflect.TypeOf((*CrossVersionObjectReferenceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CsiDriverSpec",
		reflect.TypeOf((*CsiDriverSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CsiDriverSpecV1Beta1",
		reflect.TypeOf((*CsiDriverSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CsiNodeDriver",
		reflect.TypeOf((*CsiNodeDriver)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CsiNodeDriverV1Beta1",
		reflect.TypeOf((*CsiNodeDriverV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CsiNodeSpec",
		reflect.TypeOf((*CsiNodeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CsiNodeSpecV1Beta1",
		reflect.TypeOf((*CsiNodeSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CsiPersistentVolumeSource",
		reflect.TypeOf((*CsiPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CsiVolumeSource",
		reflect.TypeOf((*CsiVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceColumnDefinition",
		reflect.TypeOf((*CustomResourceColumnDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceColumnDefinitionV1Beta1",
		reflect.TypeOf((*CustomResourceColumnDefinitionV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceConversion",
		reflect.TypeOf((*CustomResourceConversion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceConversionV1Beta1",
		reflect.TypeOf((*CustomResourceConversionV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceDefinitionNames",
		reflect.TypeOf((*CustomResourceDefinitionNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceDefinitionNamesV1Beta1",
		reflect.TypeOf((*CustomResourceDefinitionNamesV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceDefinitionSpec",
		reflect.TypeOf((*CustomResourceDefinitionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceDefinitionSpecV1Beta1",
		reflect.TypeOf((*CustomResourceDefinitionSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceDefinitionVersion",
		reflect.TypeOf((*CustomResourceDefinitionVersion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceDefinitionVersionV1Beta1",
		reflect.TypeOf((*CustomResourceDefinitionVersionV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceSubresourceScale",
		reflect.TypeOf((*CustomResourceSubresourceScale)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceSubresourceScaleV1Beta1",
		reflect.TypeOf((*CustomResourceSubresourceScaleV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceSubresources",
		reflect.TypeOf((*CustomResourceSubresources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceSubresourcesV1Beta1",
		reflect.TypeOf((*CustomResourceSubresourcesV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceValidation",
		reflect.TypeOf((*CustomResourceValidation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.CustomResourceValidationV1Beta1",
		reflect.TypeOf((*CustomResourceValidationV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.DaemonSetSpec",
		reflect.TypeOf((*DaemonSetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.DaemonSetUpdateStrategy",
		reflect.TypeOf((*DaemonSetUpdateStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.DeleteOptions",
		reflect.TypeOf((*DeleteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.DeploymentSpec",
		reflect.TypeOf((*DeploymentSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.DeploymentStrategy",
		reflect.TypeOf((*DeploymentStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.DownwardApiProjection",
		reflect.TypeOf((*DownwardApiProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.DownwardApiVolumeFile",
		reflect.TypeOf((*DownwardApiVolumeFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.DownwardApiVolumeSource",
		reflect.TypeOf((*DownwardApiVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EmptyDirVolumeSource",
		reflect.TypeOf((*EmptyDirVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Endpoint",
		reflect.TypeOf((*Endpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointAddress",
		reflect.TypeOf((*EndpointAddress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointConditions",
		reflect.TypeOf((*EndpointConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointConditionsV1Beta1",
		reflect.TypeOf((*EndpointConditionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointHints",
		reflect.TypeOf((*EndpointHints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointHintsV1Beta1",
		reflect.TypeOf((*EndpointHintsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointPort",
		reflect.TypeOf((*EndpointPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointPortV1Beta1",
		reflect.TypeOf((*EndpointPortV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointSubset",
		reflect.TypeOf((*EndpointSubset)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EndpointV1Beta1",
		reflect.TypeOf((*EndpointV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EnvFromSource",
		reflect.TypeOf((*EnvFromSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EnvVar",
		reflect.TypeOf((*EnvVar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EnvVarSource",
		reflect.TypeOf((*EnvVarSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EphemeralContainer",
		reflect.TypeOf((*EphemeralContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EphemeralVolumeSource",
		reflect.TypeOf((*EphemeralVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EventSeries",
		reflect.TypeOf((*EventSeries)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EventSeriesV1Beta1",
		reflect.TypeOf((*EventSeriesV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.EventSource",
		reflect.TypeOf((*EventSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ExecAction",
		reflect.TypeOf((*ExecAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ExternalDocumentation",
		reflect.TypeOf((*ExternalDocumentation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ExternalDocumentationV1Beta1",
		reflect.TypeOf((*ExternalDocumentationV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ExternalMetricSourceV2Beta1",
		reflect.TypeOf((*ExternalMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ExternalMetricSourceV2Beta2",
		reflect.TypeOf((*ExternalMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.FcVolumeSource",
		reflect.TypeOf((*FcVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.FlexPersistentVolumeSource",
		reflect.TypeOf((*FlexPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.FlexVolumeSource",
		reflect.TypeOf((*FlexVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.FlockerVolumeSource",
		reflect.TypeOf((*FlockerVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.FlowDistinguisherMethodV1Beta1",
		reflect.TypeOf((*FlowDistinguisherMethodV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.FlowSchemaSpecV1Beta1",
		reflect.TypeOf((*FlowSchemaSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ForZone",
		reflect.TypeOf((*ForZone)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ForZoneV1Beta1",
		reflect.TypeOf((*ForZoneV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.FsGroupStrategyOptionsV1Beta1",
		reflect.TypeOf((*FsGroupStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.GcePersistentDiskVolumeSource",
		reflect.TypeOf((*GcePersistentDiskVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.GitRepoVolumeSource",
		reflect.TypeOf((*GitRepoVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.GlusterfsPersistentVolumeSource",
		reflect.TypeOf((*GlusterfsPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.GlusterfsVolumeSource",
		reflect.TypeOf((*GlusterfsVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Handler",
		reflect.TypeOf((*Handler)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HorizontalPodAutoscalerBehaviorV2Beta2",
		reflect.TypeOf((*HorizontalPodAutoscalerBehaviorV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HorizontalPodAutoscalerSpec",
		reflect.TypeOf((*HorizontalPodAutoscalerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HorizontalPodAutoscalerSpecV2Beta1",
		reflect.TypeOf((*HorizontalPodAutoscalerSpecV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HorizontalPodAutoscalerSpecV2Beta2",
		reflect.TypeOf((*HorizontalPodAutoscalerSpecV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HostAlias",
		reflect.TypeOf((*HostAlias)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HostPathVolumeSource",
		reflect.TypeOf((*HostPathVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HostPortRangeV1Beta1",
		reflect.TypeOf((*HostPortRangeV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HpaScalingPolicyV2Beta2",
		reflect.TypeOf((*HpaScalingPolicyV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HpaScalingRulesV2Beta2",
		reflect.TypeOf((*HpaScalingRulesV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HttpGetAction",
		reflect.TypeOf((*HttpGetAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HttpHeader",
		reflect.TypeOf((*HttpHeader)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HttpIngressPath",
		reflect.TypeOf((*HttpIngressPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HttpIngressPathV1Beta1",
		reflect.TypeOf((*HttpIngressPathV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HttpIngressRuleValue",
		reflect.TypeOf((*HttpIngressRuleValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.HttpIngressRuleValueV1Beta1",
		reflect.TypeOf((*HttpIngressRuleValueV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IdRangeV1Beta1",
		reflect.TypeOf((*IdRangeV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressBackend",
		reflect.TypeOf((*IngressBackend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressBackendV1Beta1",
		reflect.TypeOf((*IngressBackendV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressClassParametersReference",
		reflect.TypeOf((*IngressClassParametersReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressClassParametersReferenceV1Beta1",
		reflect.TypeOf((*IngressClassParametersReferenceV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressClassSpec",
		reflect.TypeOf((*IngressClassSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressClassSpecV1Beta1",
		reflect.TypeOf((*IngressClassSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressRule",
		reflect.TypeOf((*IngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressRuleV1Beta1",
		reflect.TypeOf((*IngressRuleV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressServiceBackend",
		reflect.TypeOf((*IngressServiceBackend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressSpec",
		reflect.TypeOf((*IngressSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressSpecV1Beta1",
		reflect.TypeOf((*IngressSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressTls",
		reflect.TypeOf((*IngressTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IngressTlsv1Beta1",
		reflect.TypeOf((*IngressTlsv1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.IntOrString",
		reflect.TypeOf((*IntOrString)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IntOrString{}
		},
	)
	_jsii_.RegisterEnum(
		"@opencdk8s/cdk8s-external-dns-route53.IoK8SApimachineryPkgApisMetaV1DeleteOptionsKind",
		reflect.TypeOf((*IoK8SApimachineryPkgApisMetaV1DeleteOptionsKind)(nil)).Elem(),
		map[string]interface{}{
			"DELETE_OPTIONS": IoK8SApimachineryPkgApisMetaV1DeleteOptionsKind_DELETE_OPTIONS,
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IpBlock",
		reflect.TypeOf((*IpBlock)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IscsiPersistentVolumeSource",
		reflect.TypeOf((*IscsiPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.IscsiVolumeSource",
		reflect.TypeOf((*IscsiVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.JobSpec",
		reflect.TypeOf((*JobSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.JobTemplateSpec",
		reflect.TypeOf((*JobTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.JobTemplateSpecV1Beta1",
		reflect.TypeOf((*JobTemplateSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.JsonSchemaProps",
		reflect.TypeOf((*JsonSchemaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.JsonSchemaPropsV1Beta1",
		reflect.TypeOf((*JsonSchemaPropsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KeyToPath",
		reflect.TypeOf((*KeyToPath)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeApiService",
		reflect.TypeOf((*KubeApiService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeApiService{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeApiServiceList",
		reflect.TypeOf((*KubeApiServiceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeApiServiceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeApiServiceListProps",
		reflect.TypeOf((*KubeApiServiceListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeApiServiceListV1Beta1",
		reflect.TypeOf((*KubeApiServiceListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeApiServiceListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeApiServiceListV1Beta1Props",
		reflect.TypeOf((*KubeApiServiceListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeApiServiceProps",
		reflect.TypeOf((*KubeApiServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeApiServiceV1Beta1",
		reflect.TypeOf((*KubeApiServiceV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeApiServiceV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeApiServiceV1Beta1Props",
		reflect.TypeOf((*KubeApiServiceV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeBinding",
		reflect.TypeOf((*KubeBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeBinding{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeBindingProps",
		reflect.TypeOf((*KubeBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCertificateSigningRequest",
		reflect.TypeOf((*KubeCertificateSigningRequest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCertificateSigningRequest{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCertificateSigningRequestList",
		reflect.TypeOf((*KubeCertificateSigningRequestList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCertificateSigningRequestList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCertificateSigningRequestListProps",
		reflect.TypeOf((*KubeCertificateSigningRequestListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCertificateSigningRequestListV1Beta1",
		reflect.TypeOf((*KubeCertificateSigningRequestListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCertificateSigningRequestListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCertificateSigningRequestListV1Beta1Props",
		reflect.TypeOf((*KubeCertificateSigningRequestListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCertificateSigningRequestProps",
		reflect.TypeOf((*KubeCertificateSigningRequestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCertificateSigningRequestV1Beta1",
		reflect.TypeOf((*KubeCertificateSigningRequestV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCertificateSigningRequestV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCertificateSigningRequestV1Beta1Props",
		reflect.TypeOf((*KubeCertificateSigningRequestV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRole",
		reflect.TypeOf((*KubeClusterRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBinding",
		reflect.TypeOf((*KubeClusterRoleBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleBinding{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingList",
		reflect.TypeOf((*KubeClusterRoleBindingList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleBindingList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingListProps",
		reflect.TypeOf((*KubeClusterRoleBindingListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingListV1Alpha1",
		reflect.TypeOf((*KubeClusterRoleBindingListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleBindingListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingListV1Alpha1Props",
		reflect.TypeOf((*KubeClusterRoleBindingListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingListV1Beta1",
		reflect.TypeOf((*KubeClusterRoleBindingListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleBindingListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingListV1Beta1Props",
		reflect.TypeOf((*KubeClusterRoleBindingListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingProps",
		reflect.TypeOf((*KubeClusterRoleBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingV1Alpha1",
		reflect.TypeOf((*KubeClusterRoleBindingV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleBindingV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingV1Alpha1Props",
		reflect.TypeOf((*KubeClusterRoleBindingV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingV1Beta1",
		reflect.TypeOf((*KubeClusterRoleBindingV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleBindingV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleBindingV1Beta1Props",
		reflect.TypeOf((*KubeClusterRoleBindingV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleList",
		reflect.TypeOf((*KubeClusterRoleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleListProps",
		reflect.TypeOf((*KubeClusterRoleListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleListV1Alpha1",
		reflect.TypeOf((*KubeClusterRoleListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleListV1Alpha1Props",
		reflect.TypeOf((*KubeClusterRoleListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleListV1Beta1",
		reflect.TypeOf((*KubeClusterRoleListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleListV1Beta1Props",
		reflect.TypeOf((*KubeClusterRoleListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleProps",
		reflect.TypeOf((*KubeClusterRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleV1Alpha1",
		reflect.TypeOf((*KubeClusterRoleV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleV1Alpha1Props",
		reflect.TypeOf((*KubeClusterRoleV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleV1Beta1",
		reflect.TypeOf((*KubeClusterRoleV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeClusterRoleV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeClusterRoleV1Beta1Props",
		reflect.TypeOf((*KubeClusterRoleV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeComponentStatus",
		reflect.TypeOf((*KubeComponentStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeComponentStatus{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeComponentStatusList",
		reflect.TypeOf((*KubeComponentStatusList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeComponentStatusList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeComponentStatusListProps",
		reflect.TypeOf((*KubeComponentStatusListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeComponentStatusProps",
		reflect.TypeOf((*KubeComponentStatusProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeConfigMap",
		reflect.TypeOf((*KubeConfigMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeConfigMap{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeConfigMapList",
		reflect.TypeOf((*KubeConfigMapList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeConfigMapList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeConfigMapListProps",
		reflect.TypeOf((*KubeConfigMapListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeConfigMapProps",
		reflect.TypeOf((*KubeConfigMapProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeControllerRevision",
		reflect.TypeOf((*KubeControllerRevision)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeControllerRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeControllerRevisionList",
		reflect.TypeOf((*KubeControllerRevisionList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeControllerRevisionList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeControllerRevisionListProps",
		reflect.TypeOf((*KubeControllerRevisionListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeControllerRevisionProps",
		reflect.TypeOf((*KubeControllerRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCronJob",
		reflect.TypeOf((*KubeCronJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCronJob{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCronJobList",
		reflect.TypeOf((*KubeCronJobList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCronJobList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCronJobListProps",
		reflect.TypeOf((*KubeCronJobListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCronJobListV1Beta1",
		reflect.TypeOf((*KubeCronJobListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCronJobListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCronJobListV1Beta1Props",
		reflect.TypeOf((*KubeCronJobListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCronJobProps",
		reflect.TypeOf((*KubeCronJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCronJobV1Beta1",
		reflect.TypeOf((*KubeCronJobV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCronJobV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCronJobV1Beta1Props",
		reflect.TypeOf((*KubeCronJobV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiDriver",
		reflect.TypeOf((*KubeCsiDriver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiDriver{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiDriverList",
		reflect.TypeOf((*KubeCsiDriverList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiDriverList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiDriverListProps",
		reflect.TypeOf((*KubeCsiDriverListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiDriverListV1Beta1",
		reflect.TypeOf((*KubeCsiDriverListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiDriverListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiDriverListV1Beta1Props",
		reflect.TypeOf((*KubeCsiDriverListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiDriverProps",
		reflect.TypeOf((*KubeCsiDriverProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiDriverV1Beta1",
		reflect.TypeOf((*KubeCsiDriverV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiDriverV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiDriverV1Beta1Props",
		reflect.TypeOf((*KubeCsiDriverV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiNode",
		reflect.TypeOf((*KubeCsiNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiNode{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiNodeList",
		reflect.TypeOf((*KubeCsiNodeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiNodeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiNodeListProps",
		reflect.TypeOf((*KubeCsiNodeListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiNodeListV1Beta1",
		reflect.TypeOf((*KubeCsiNodeListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiNodeListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiNodeListV1Beta1Props",
		reflect.TypeOf((*KubeCsiNodeListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiNodeProps",
		reflect.TypeOf((*KubeCsiNodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiNodeV1Beta1",
		reflect.TypeOf((*KubeCsiNodeV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiNodeV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiNodeV1Beta1Props",
		reflect.TypeOf((*KubeCsiNodeV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiStorageCapacityListV1Alpha1",
		reflect.TypeOf((*KubeCsiStorageCapacityListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiStorageCapacityListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiStorageCapacityListV1Alpha1Props",
		reflect.TypeOf((*KubeCsiStorageCapacityListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiStorageCapacityListV1Beta1",
		reflect.TypeOf((*KubeCsiStorageCapacityListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiStorageCapacityListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiStorageCapacityListV1Beta1Props",
		reflect.TypeOf((*KubeCsiStorageCapacityListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiStorageCapacityV1Alpha1",
		reflect.TypeOf((*KubeCsiStorageCapacityV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiStorageCapacityV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiStorageCapacityV1Alpha1Props",
		reflect.TypeOf((*KubeCsiStorageCapacityV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiStorageCapacityV1Beta1",
		reflect.TypeOf((*KubeCsiStorageCapacityV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCsiStorageCapacityV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCsiStorageCapacityV1Beta1Props",
		reflect.TypeOf((*KubeCsiStorageCapacityV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCustomResourceDefinition",
		reflect.TypeOf((*KubeCustomResourceDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCustomResourceDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCustomResourceDefinitionList",
		reflect.TypeOf((*KubeCustomResourceDefinitionList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCustomResourceDefinitionList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCustomResourceDefinitionListProps",
		reflect.TypeOf((*KubeCustomResourceDefinitionListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCustomResourceDefinitionListV1Beta1",
		reflect.TypeOf((*KubeCustomResourceDefinitionListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCustomResourceDefinitionListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCustomResourceDefinitionListV1Beta1Props",
		reflect.TypeOf((*KubeCustomResourceDefinitionListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCustomResourceDefinitionProps",
		reflect.TypeOf((*KubeCustomResourceDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCustomResourceDefinitionV1Beta1",
		reflect.TypeOf((*KubeCustomResourceDefinitionV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeCustomResourceDefinitionV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeCustomResourceDefinitionV1Beta1Props",
		reflect.TypeOf((*KubeCustomResourceDefinitionV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeDaemonSet",
		reflect.TypeOf((*KubeDaemonSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeDaemonSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeDaemonSetList",
		reflect.TypeOf((*KubeDaemonSetList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeDaemonSetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeDaemonSetListProps",
		reflect.TypeOf((*KubeDaemonSetListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeDaemonSetProps",
		reflect.TypeOf((*KubeDaemonSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeDeployment",
		reflect.TypeOf((*KubeDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeDeploymentList",
		reflect.TypeOf((*KubeDeploymentList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeDeploymentList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeDeploymentListProps",
		reflect.TypeOf((*KubeDeploymentListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeDeploymentProps",
		reflect.TypeOf((*KubeDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointSlice",
		reflect.TypeOf((*KubeEndpointSlice)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEndpointSlice{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointSliceList",
		reflect.TypeOf((*KubeEndpointSliceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEndpointSliceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointSliceListProps",
		reflect.TypeOf((*KubeEndpointSliceListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointSliceListV1Beta1",
		reflect.TypeOf((*KubeEndpointSliceListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEndpointSliceListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointSliceListV1Beta1Props",
		reflect.TypeOf((*KubeEndpointSliceListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointSliceProps",
		reflect.TypeOf((*KubeEndpointSliceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointSliceV1Beta1",
		reflect.TypeOf((*KubeEndpointSliceV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEndpointSliceV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointSliceV1Beta1Props",
		reflect.TypeOf((*KubeEndpointSliceV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpoints",
		reflect.TypeOf((*KubeEndpoints)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEndpoints{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointsList",
		reflect.TypeOf((*KubeEndpointsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEndpointsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointsListProps",
		reflect.TypeOf((*KubeEndpointsListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEndpointsProps",
		reflect.TypeOf((*KubeEndpointsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEphemeralContainers",
		reflect.TypeOf((*KubeEphemeralContainers)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEphemeralContainers{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEphemeralContainersProps",
		reflect.TypeOf((*KubeEphemeralContainersProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEvent",
		reflect.TypeOf((*KubeEvent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEvent{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEventList",
		reflect.TypeOf((*KubeEventList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEventList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEventListProps",
		reflect.TypeOf((*KubeEventListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEventListV1Beta1",
		reflect.TypeOf((*KubeEventListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEventListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEventListV1Beta1Props",
		reflect.TypeOf((*KubeEventListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEventProps",
		reflect.TypeOf((*KubeEventProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEventV1Beta1",
		reflect.TypeOf((*KubeEventV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEventV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEventV1Beta1Props",
		reflect.TypeOf((*KubeEventV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEvictionV1Beta1",
		reflect.TypeOf((*KubeEvictionV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeEvictionV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeEvictionV1Beta1Props",
		reflect.TypeOf((*KubeEvictionV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeFlowSchemaListV1Beta1",
		reflect.TypeOf((*KubeFlowSchemaListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeFlowSchemaListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeFlowSchemaListV1Beta1Props",
		reflect.TypeOf((*KubeFlowSchemaListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeFlowSchemaV1Beta1",
		reflect.TypeOf((*KubeFlowSchemaV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeFlowSchemaV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeFlowSchemaV1Beta1Props",
		reflect.TypeOf((*KubeFlowSchemaV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscaler",
		reflect.TypeOf((*KubeHorizontalPodAutoscaler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeHorizontalPodAutoscaler{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerList",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeHorizontalPodAutoscalerList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerListProps",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerListV2Beta1",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerListV2Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeHorizontalPodAutoscalerListV2Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerListV2Beta1Props",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerListV2Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerListV2Beta2",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerListV2Beta2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeHorizontalPodAutoscalerListV2Beta2{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerListV2Beta2Props",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerListV2Beta2Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerProps",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerV2Beta1",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerV2Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeHorizontalPodAutoscalerV2Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerV2Beta1Props",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerV2Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerV2Beta2",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerV2Beta2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeHorizontalPodAutoscalerV2Beta2{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeHorizontalPodAutoscalerV2Beta2Props",
		reflect.TypeOf((*KubeHorizontalPodAutoscalerV2Beta2Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngress",
		reflect.TypeOf((*KubeIngress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeIngress{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressClass",
		reflect.TypeOf((*KubeIngressClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeIngressClass{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressClassList",
		reflect.TypeOf((*KubeIngressClassList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeIngressClassList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressClassListProps",
		reflect.TypeOf((*KubeIngressClassListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressClassListV1Beta1",
		reflect.TypeOf((*KubeIngressClassListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeIngressClassListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressClassListV1Beta1Props",
		reflect.TypeOf((*KubeIngressClassListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressClassProps",
		reflect.TypeOf((*KubeIngressClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressClassV1Beta1",
		reflect.TypeOf((*KubeIngressClassV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeIngressClassV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressClassV1Beta1Props",
		reflect.TypeOf((*KubeIngressClassV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressList",
		reflect.TypeOf((*KubeIngressList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeIngressList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressListProps",
		reflect.TypeOf((*KubeIngressListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressListV1Beta1",
		reflect.TypeOf((*KubeIngressListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeIngressListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressListV1Beta1Props",
		reflect.TypeOf((*KubeIngressListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressProps",
		reflect.TypeOf((*KubeIngressProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressV1Beta1",
		reflect.TypeOf((*KubeIngressV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeIngressV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeIngressV1Beta1Props",
		reflect.TypeOf((*KubeIngressV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeJob",
		reflect.TypeOf((*KubeJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeJob{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeJobList",
		reflect.TypeOf((*KubeJobList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeJobList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeJobListProps",
		reflect.TypeOf((*KubeJobListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeJobProps",
		reflect.TypeOf((*KubeJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLease",
		reflect.TypeOf((*KubeLease)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeLease{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLeaseList",
		reflect.TypeOf((*KubeLeaseList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeLeaseList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLeaseListProps",
		reflect.TypeOf((*KubeLeaseListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLeaseListV1Beta1",
		reflect.TypeOf((*KubeLeaseListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeLeaseListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLeaseListV1Beta1Props",
		reflect.TypeOf((*KubeLeaseListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLeaseProps",
		reflect.TypeOf((*KubeLeaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLeaseV1Beta1",
		reflect.TypeOf((*KubeLeaseV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeLeaseV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLeaseV1Beta1Props",
		reflect.TypeOf((*KubeLeaseV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLimitRange",
		reflect.TypeOf((*KubeLimitRange)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeLimitRange{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLimitRangeList",
		reflect.TypeOf((*KubeLimitRangeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeLimitRangeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLimitRangeListProps",
		reflect.TypeOf((*KubeLimitRangeListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLimitRangeProps",
		reflect.TypeOf((*KubeLimitRangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLocalSubjectAccessReview",
		reflect.TypeOf((*KubeLocalSubjectAccessReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeLocalSubjectAccessReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLocalSubjectAccessReviewProps",
		reflect.TypeOf((*KubeLocalSubjectAccessReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLocalSubjectAccessReviewV1Beta1",
		reflect.TypeOf((*KubeLocalSubjectAccessReviewV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeLocalSubjectAccessReviewV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeLocalSubjectAccessReviewV1Beta1Props",
		reflect.TypeOf((*KubeLocalSubjectAccessReviewV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeMutatingWebhookConfiguration",
		reflect.TypeOf((*KubeMutatingWebhookConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeMutatingWebhookConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeMutatingWebhookConfigurationList",
		reflect.TypeOf((*KubeMutatingWebhookConfigurationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeMutatingWebhookConfigurationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeMutatingWebhookConfigurationListProps",
		reflect.TypeOf((*KubeMutatingWebhookConfigurationListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeMutatingWebhookConfigurationListV1Beta1",
		reflect.TypeOf((*KubeMutatingWebhookConfigurationListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeMutatingWebhookConfigurationListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeMutatingWebhookConfigurationListV1Beta1Props",
		reflect.TypeOf((*KubeMutatingWebhookConfigurationListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeMutatingWebhookConfigurationProps",
		reflect.TypeOf((*KubeMutatingWebhookConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeMutatingWebhookConfigurationV1Beta1",
		reflect.TypeOf((*KubeMutatingWebhookConfigurationV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeMutatingWebhookConfigurationV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeMutatingWebhookConfigurationV1Beta1Props",
		reflect.TypeOf((*KubeMutatingWebhookConfigurationV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNamespace",
		reflect.TypeOf((*KubeNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNamespaceList",
		reflect.TypeOf((*KubeNamespaceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeNamespaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNamespaceListProps",
		reflect.TypeOf((*KubeNamespaceListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNamespaceProps",
		reflect.TypeOf((*KubeNamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNetworkPolicy",
		reflect.TypeOf((*KubeNetworkPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeNetworkPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNetworkPolicyList",
		reflect.TypeOf((*KubeNetworkPolicyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeNetworkPolicyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNetworkPolicyListProps",
		reflect.TypeOf((*KubeNetworkPolicyListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNetworkPolicyProps",
		reflect.TypeOf((*KubeNetworkPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNode",
		reflect.TypeOf((*KubeNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeNode{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNodeList",
		reflect.TypeOf((*KubeNodeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeNodeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNodeListProps",
		reflect.TypeOf((*KubeNodeListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeNodeProps",
		reflect.TypeOf((*KubeNodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePersistentVolume",
		reflect.TypeOf((*KubePersistentVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePersistentVolume{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePersistentVolumeClaim",
		reflect.TypeOf((*KubePersistentVolumeClaim)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePersistentVolumeClaim{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePersistentVolumeClaimList",
		reflect.TypeOf((*KubePersistentVolumeClaimList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePersistentVolumeClaimList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePersistentVolumeClaimListProps",
		reflect.TypeOf((*KubePersistentVolumeClaimListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePersistentVolumeClaimProps",
		reflect.TypeOf((*KubePersistentVolumeClaimProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePersistentVolumeList",
		reflect.TypeOf((*KubePersistentVolumeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePersistentVolumeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePersistentVolumeListProps",
		reflect.TypeOf((*KubePersistentVolumeListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePersistentVolumeProps",
		reflect.TypeOf((*KubePersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePod",
		reflect.TypeOf((*KubePod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePod{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodDisruptionBudget",
		reflect.TypeOf((*KubePodDisruptionBudget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodDisruptionBudget{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodDisruptionBudgetList",
		reflect.TypeOf((*KubePodDisruptionBudgetList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodDisruptionBudgetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodDisruptionBudgetListProps",
		reflect.TypeOf((*KubePodDisruptionBudgetListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodDisruptionBudgetListV1Beta1",
		reflect.TypeOf((*KubePodDisruptionBudgetListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodDisruptionBudgetListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodDisruptionBudgetListV1Beta1Props",
		reflect.TypeOf((*KubePodDisruptionBudgetListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodDisruptionBudgetProps",
		reflect.TypeOf((*KubePodDisruptionBudgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodDisruptionBudgetV1Beta1",
		reflect.TypeOf((*KubePodDisruptionBudgetV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodDisruptionBudgetV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodDisruptionBudgetV1Beta1Props",
		reflect.TypeOf((*KubePodDisruptionBudgetV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodList",
		reflect.TypeOf((*KubePodList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodListProps",
		reflect.TypeOf((*KubePodListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodProps",
		reflect.TypeOf((*KubePodProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodSecurityPolicyListV1Beta1",
		reflect.TypeOf((*KubePodSecurityPolicyListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodSecurityPolicyListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodSecurityPolicyListV1Beta1Props",
		reflect.TypeOf((*KubePodSecurityPolicyListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodSecurityPolicyV1Beta1",
		reflect.TypeOf((*KubePodSecurityPolicyV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodSecurityPolicyV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodSecurityPolicyV1Beta1Props",
		reflect.TypeOf((*KubePodSecurityPolicyV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodTemplate",
		reflect.TypeOf((*KubePodTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodTemplateList",
		reflect.TypeOf((*KubePodTemplateList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePodTemplateList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodTemplateListProps",
		reflect.TypeOf((*KubePodTemplateListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePodTemplateProps",
		reflect.TypeOf((*KubePodTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClass",
		reflect.TypeOf((*KubePriorityClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePriorityClass{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassList",
		reflect.TypeOf((*KubePriorityClassList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePriorityClassList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassListProps",
		reflect.TypeOf((*KubePriorityClassListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassListV1Alpha1",
		reflect.TypeOf((*KubePriorityClassListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePriorityClassListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassListV1Alpha1Props",
		reflect.TypeOf((*KubePriorityClassListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassListV1Beta1",
		reflect.TypeOf((*KubePriorityClassListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePriorityClassListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassListV1Beta1Props",
		reflect.TypeOf((*KubePriorityClassListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassProps",
		reflect.TypeOf((*KubePriorityClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassV1Alpha1",
		reflect.TypeOf((*KubePriorityClassV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePriorityClassV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassV1Alpha1Props",
		reflect.TypeOf((*KubePriorityClassV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassV1Beta1",
		reflect.TypeOf((*KubePriorityClassV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePriorityClassV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityClassV1Beta1Props",
		reflect.TypeOf((*KubePriorityClassV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityLevelConfigurationListV1Beta1",
		reflect.TypeOf((*KubePriorityLevelConfigurationListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePriorityLevelConfigurationListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityLevelConfigurationListV1Beta1Props",
		reflect.TypeOf((*KubePriorityLevelConfigurationListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityLevelConfigurationV1Beta1",
		reflect.TypeOf((*KubePriorityLevelConfigurationV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubePriorityLevelConfigurationV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubePriorityLevelConfigurationV1Beta1Props",
		reflect.TypeOf((*KubePriorityLevelConfigurationV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeReplicaSet",
		reflect.TypeOf((*KubeReplicaSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeReplicaSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeReplicaSetList",
		reflect.TypeOf((*KubeReplicaSetList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeReplicaSetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeReplicaSetListProps",
		reflect.TypeOf((*KubeReplicaSetListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeReplicaSetProps",
		reflect.TypeOf((*KubeReplicaSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeReplicationController",
		reflect.TypeOf((*KubeReplicationController)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeReplicationController{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeReplicationControllerList",
		reflect.TypeOf((*KubeReplicationControllerList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeReplicationControllerList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeReplicationControllerListProps",
		reflect.TypeOf((*KubeReplicationControllerListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeReplicationControllerProps",
		reflect.TypeOf((*KubeReplicationControllerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeResourceQuota",
		reflect.TypeOf((*KubeResourceQuota)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeResourceQuota{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeResourceQuotaList",
		reflect.TypeOf((*KubeResourceQuotaList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeResourceQuotaList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeResourceQuotaListProps",
		reflect.TypeOf((*KubeResourceQuotaListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeResourceQuotaProps",
		reflect.TypeOf((*KubeResourceQuotaProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRole",
		reflect.TypeOf((*KubeRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBinding",
		reflect.TypeOf((*KubeRoleBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleBinding{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingList",
		reflect.TypeOf((*KubeRoleBindingList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleBindingList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingListProps",
		reflect.TypeOf((*KubeRoleBindingListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingListV1Alpha1",
		reflect.TypeOf((*KubeRoleBindingListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleBindingListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingListV1Alpha1Props",
		reflect.TypeOf((*KubeRoleBindingListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingListV1Beta1",
		reflect.TypeOf((*KubeRoleBindingListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleBindingListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingListV1Beta1Props",
		reflect.TypeOf((*KubeRoleBindingListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingProps",
		reflect.TypeOf((*KubeRoleBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingV1Alpha1",
		reflect.TypeOf((*KubeRoleBindingV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleBindingV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingV1Alpha1Props",
		reflect.TypeOf((*KubeRoleBindingV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingV1Beta1",
		reflect.TypeOf((*KubeRoleBindingV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleBindingV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleBindingV1Beta1Props",
		reflect.TypeOf((*KubeRoleBindingV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleList",
		reflect.TypeOf((*KubeRoleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleListProps",
		reflect.TypeOf((*KubeRoleListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleListV1Alpha1",
		reflect.TypeOf((*KubeRoleListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleListV1Alpha1Props",
		reflect.TypeOf((*KubeRoleListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleListV1Beta1",
		reflect.TypeOf((*KubeRoleListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleListV1Beta1Props",
		reflect.TypeOf((*KubeRoleListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleProps",
		reflect.TypeOf((*KubeRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleV1Alpha1",
		reflect.TypeOf((*KubeRoleV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleV1Alpha1Props",
		reflect.TypeOf((*KubeRoleV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleV1Beta1",
		reflect.TypeOf((*KubeRoleV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRoleV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRoleV1Beta1Props",
		reflect.TypeOf((*KubeRoleV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClass",
		reflect.TypeOf((*KubeRuntimeClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRuntimeClass{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassList",
		reflect.TypeOf((*KubeRuntimeClassList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRuntimeClassList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassListProps",
		reflect.TypeOf((*KubeRuntimeClassListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassListV1Alpha1",
		reflect.TypeOf((*KubeRuntimeClassListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRuntimeClassListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassListV1Alpha1Props",
		reflect.TypeOf((*KubeRuntimeClassListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassListV1Beta1",
		reflect.TypeOf((*KubeRuntimeClassListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRuntimeClassListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassListV1Beta1Props",
		reflect.TypeOf((*KubeRuntimeClassListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassProps",
		reflect.TypeOf((*KubeRuntimeClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassV1Alpha1",
		reflect.TypeOf((*KubeRuntimeClassV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRuntimeClassV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassV1Alpha1Props",
		reflect.TypeOf((*KubeRuntimeClassV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassV1Beta1",
		reflect.TypeOf((*KubeRuntimeClassV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeRuntimeClassV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeRuntimeClassV1Beta1Props",
		reflect.TypeOf((*KubeRuntimeClassV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeScale",
		reflect.TypeOf((*KubeScale)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeScale{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeScaleProps",
		reflect.TypeOf((*KubeScaleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSecret",
		reflect.TypeOf((*KubeSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeSecret{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSecretList",
		reflect.TypeOf((*KubeSecretList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeSecretList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSecretListProps",
		reflect.TypeOf((*KubeSecretListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSecretProps",
		reflect.TypeOf((*KubeSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSelfSubjectAccessReview",
		reflect.TypeOf((*KubeSelfSubjectAccessReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeSelfSubjectAccessReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSelfSubjectAccessReviewProps",
		reflect.TypeOf((*KubeSelfSubjectAccessReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSelfSubjectAccessReviewV1Beta1",
		reflect.TypeOf((*KubeSelfSubjectAccessReviewV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeSelfSubjectAccessReviewV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSelfSubjectAccessReviewV1Beta1Props",
		reflect.TypeOf((*KubeSelfSubjectAccessReviewV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSelfSubjectRulesReview",
		reflect.TypeOf((*KubeSelfSubjectRulesReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeSelfSubjectRulesReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSelfSubjectRulesReviewProps",
		reflect.TypeOf((*KubeSelfSubjectRulesReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSelfSubjectRulesReviewV1Beta1",
		reflect.TypeOf((*KubeSelfSubjectRulesReviewV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeSelfSubjectRulesReviewV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSelfSubjectRulesReviewV1Beta1Props",
		reflect.TypeOf((*KubeSelfSubjectRulesReviewV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeService",
		reflect.TypeOf((*KubeService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeService{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeServiceAccount",
		reflect.TypeOf((*KubeServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeServiceAccount{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeServiceAccountList",
		reflect.TypeOf((*KubeServiceAccountList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeServiceAccountList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeServiceAccountListProps",
		reflect.TypeOf((*KubeServiceAccountListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeServiceAccountProps",
		reflect.TypeOf((*KubeServiceAccountProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeServiceList",
		reflect.TypeOf((*KubeServiceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeServiceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeServiceListProps",
		reflect.TypeOf((*KubeServiceListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeServiceProps",
		reflect.TypeOf((*KubeServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStatefulSet",
		reflect.TypeOf((*KubeStatefulSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStatefulSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStatefulSetList",
		reflect.TypeOf((*KubeStatefulSetList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStatefulSetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStatefulSetListProps",
		reflect.TypeOf((*KubeStatefulSetListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStatefulSetProps",
		reflect.TypeOf((*KubeStatefulSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStatus",
		reflect.TypeOf((*KubeStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStatus{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStatusProps",
		reflect.TypeOf((*KubeStatusProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageClass",
		reflect.TypeOf((*KubeStorageClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStorageClass{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageClassList",
		reflect.TypeOf((*KubeStorageClassList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStorageClassList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageClassListProps",
		reflect.TypeOf((*KubeStorageClassListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageClassListV1Beta1",
		reflect.TypeOf((*KubeStorageClassListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStorageClassListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageClassListV1Beta1Props",
		reflect.TypeOf((*KubeStorageClassListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageClassProps",
		reflect.TypeOf((*KubeStorageClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageClassV1Beta1",
		reflect.TypeOf((*KubeStorageClassV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStorageClassV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageClassV1Beta1Props",
		reflect.TypeOf((*KubeStorageClassV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageVersionListV1Alpha1",
		reflect.TypeOf((*KubeStorageVersionListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStorageVersionListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageVersionListV1Alpha1Props",
		reflect.TypeOf((*KubeStorageVersionListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageVersionV1Alpha1",
		reflect.TypeOf((*KubeStorageVersionV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeStorageVersionV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeStorageVersionV1Alpha1Props",
		reflect.TypeOf((*KubeStorageVersionV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSubjectAccessReview",
		reflect.TypeOf((*KubeSubjectAccessReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeSubjectAccessReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSubjectAccessReviewProps",
		reflect.TypeOf((*KubeSubjectAccessReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSubjectAccessReviewV1Beta1",
		reflect.TypeOf((*KubeSubjectAccessReviewV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeSubjectAccessReviewV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeSubjectAccessReviewV1Beta1Props",
		reflect.TypeOf((*KubeSubjectAccessReviewV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeTokenRequest",
		reflect.TypeOf((*KubeTokenRequest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeTokenRequest{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeTokenRequestProps",
		reflect.TypeOf((*KubeTokenRequestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeTokenReview",
		reflect.TypeOf((*KubeTokenReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeTokenReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeTokenReviewProps",
		reflect.TypeOf((*KubeTokenReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeTokenReviewV1Beta1",
		reflect.TypeOf((*KubeTokenReviewV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeTokenReviewV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeTokenReviewV1Beta1Props",
		reflect.TypeOf((*KubeTokenReviewV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeValidatingWebhookConfiguration",
		reflect.TypeOf((*KubeValidatingWebhookConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeValidatingWebhookConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeValidatingWebhookConfigurationList",
		reflect.TypeOf((*KubeValidatingWebhookConfigurationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeValidatingWebhookConfigurationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeValidatingWebhookConfigurationListProps",
		reflect.TypeOf((*KubeValidatingWebhookConfigurationListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeValidatingWebhookConfigurationListV1Beta1",
		reflect.TypeOf((*KubeValidatingWebhookConfigurationListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeValidatingWebhookConfigurationListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeValidatingWebhookConfigurationListV1Beta1Props",
		reflect.TypeOf((*KubeValidatingWebhookConfigurationListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeValidatingWebhookConfigurationProps",
		reflect.TypeOf((*KubeValidatingWebhookConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeValidatingWebhookConfigurationV1Beta1",
		reflect.TypeOf((*KubeValidatingWebhookConfigurationV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeValidatingWebhookConfigurationV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeValidatingWebhookConfigurationV1Beta1Props",
		reflect.TypeOf((*KubeValidatingWebhookConfigurationV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachment",
		reflect.TypeOf((*KubeVolumeAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeVolumeAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentList",
		reflect.TypeOf((*KubeVolumeAttachmentList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeVolumeAttachmentList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentListProps",
		reflect.TypeOf((*KubeVolumeAttachmentListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentListV1Alpha1",
		reflect.TypeOf((*KubeVolumeAttachmentListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeVolumeAttachmentListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentListV1Alpha1Props",
		reflect.TypeOf((*KubeVolumeAttachmentListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentListV1Beta1",
		reflect.TypeOf((*KubeVolumeAttachmentListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeVolumeAttachmentListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentListV1Beta1Props",
		reflect.TypeOf((*KubeVolumeAttachmentListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentProps",
		reflect.TypeOf((*KubeVolumeAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentV1Alpha1",
		reflect.TypeOf((*KubeVolumeAttachmentV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeVolumeAttachmentV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentV1Alpha1Props",
		reflect.TypeOf((*KubeVolumeAttachmentV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentV1Beta1",
		reflect.TypeOf((*KubeVolumeAttachmentV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeVolumeAttachmentV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.KubeVolumeAttachmentV1Beta1Props",
		reflect.TypeOf((*KubeVolumeAttachmentV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LabelSelector",
		reflect.TypeOf((*LabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LabelSelectorRequirement",
		reflect.TypeOf((*LabelSelectorRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LeaseSpec",
		reflect.TypeOf((*LeaseSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LeaseSpecV1Beta1",
		reflect.TypeOf((*LeaseSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Lifecycle",
		reflect.TypeOf((*Lifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LimitRangeItem",
		reflect.TypeOf((*LimitRangeItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LimitRangeSpec",
		reflect.TypeOf((*LimitRangeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LimitResponseV1Beta1",
		reflect.TypeOf((*LimitResponseV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LimitedPriorityLevelConfigurationV1Beta1",
		reflect.TypeOf((*LimitedPriorityLevelConfigurationV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ListMeta",
		reflect.TypeOf((*ListMeta)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LocalObjectReference",
		reflect.TypeOf((*LocalObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.LocalVolumeSource",
		reflect.TypeOf((*LocalVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ManagedFieldsEntry",
		reflect.TypeOf((*ManagedFieldsEntry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.MetricIdentifierV2Beta2",
		reflect.TypeOf((*MetricIdentifierV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.MetricSpecV2Beta1",
		reflect.TypeOf((*MetricSpecV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.MetricSpecV2Beta2",
		reflect.TypeOf((*MetricSpecV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.MetricTargetV2Beta2",
		reflect.TypeOf((*MetricTargetV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.MutatingWebhook",
		reflect.TypeOf((*MutatingWebhook)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.MutatingWebhookV1Beta1",
		reflect.TypeOf((*MutatingWebhookV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NamespaceSpec",
		reflect.TypeOf((*NamespaceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NetworkPolicyEgressRule",
		reflect.TypeOf((*NetworkPolicyEgressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NetworkPolicyIngressRule",
		reflect.TypeOf((*NetworkPolicyIngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NetworkPolicyPeer",
		reflect.TypeOf((*NetworkPolicyPeer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NetworkPolicyPort",
		reflect.TypeOf((*NetworkPolicyPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NetworkPolicySpec",
		reflect.TypeOf((*NetworkPolicySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NfsVolumeSource",
		reflect.TypeOf((*NfsVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NodeAffinity",
		reflect.TypeOf((*NodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NodeConfigSource",
		reflect.TypeOf((*NodeConfigSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NodeSelector",
		reflect.TypeOf((*NodeSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NodeSelectorRequirement",
		reflect.TypeOf((*NodeSelectorRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NodeSelectorTerm",
		reflect.TypeOf((*NodeSelectorTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NodeSpec",
		reflect.TypeOf((*NodeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NonResourceAttributes",
		reflect.TypeOf((*NonResourceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NonResourceAttributesV1Beta1",
		reflect.TypeOf((*NonResourceAttributesV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.NonResourcePolicyRuleV1Beta1",
		reflect.TypeOf((*NonResourcePolicyRuleV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ObjectFieldSelector",
		reflect.TypeOf((*ObjectFieldSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ObjectMeta",
		reflect.TypeOf((*ObjectMeta)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ObjectMetricSourceV2Beta1",
		reflect.TypeOf((*ObjectMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ObjectMetricSourceV2Beta2",
		reflect.TypeOf((*ObjectMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ObjectReference",
		reflect.TypeOf((*ObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Overhead",
		reflect.TypeOf((*Overhead)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.OverheadV1Alpha1",
		reflect.TypeOf((*OverheadV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.OverheadV1Beta1",
		reflect.TypeOf((*OverheadV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.OwnerReference",
		reflect.TypeOf((*OwnerReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PersistentVolumeClaimSpec",
		reflect.TypeOf((*PersistentVolumeClaimSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PersistentVolumeClaimTemplate",
		reflect.TypeOf((*PersistentVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PersistentVolumeClaimVolumeSource",
		reflect.TypeOf((*PersistentVolumeClaimVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PersistentVolumeSpec",
		reflect.TypeOf((*PersistentVolumeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PhotonPersistentDiskVolumeSource",
		reflect.TypeOf((*PhotonPersistentDiskVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodAffinity",
		reflect.TypeOf((*PodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodAffinityTerm",
		reflect.TypeOf((*PodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodAntiAffinity",
		reflect.TypeOf((*PodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodDisruptionBudgetSpec",
		reflect.TypeOf((*PodDisruptionBudgetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodDisruptionBudgetSpecV1Beta1",
		reflect.TypeOf((*PodDisruptionBudgetSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodDnsConfig",
		reflect.TypeOf((*PodDnsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodDnsConfigOption",
		reflect.TypeOf((*PodDnsConfigOption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodReadinessGate",
		reflect.TypeOf((*PodReadinessGate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodSecurityContext",
		reflect.TypeOf((*PodSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodSecurityPolicySpecV1Beta1",
		reflect.TypeOf((*PodSecurityPolicySpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodSpec",
		reflect.TypeOf((*PodSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodTemplateSpec",
		reflect.TypeOf((*PodTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodsMetricSourceV2Beta1",
		reflect.TypeOf((*PodsMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PodsMetricSourceV2Beta2",
		reflect.TypeOf((*PodsMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PolicyRule",
		reflect.TypeOf((*PolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PolicyRuleV1Alpha1",
		reflect.TypeOf((*PolicyRuleV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PolicyRuleV1Beta1",
		reflect.TypeOf((*PolicyRuleV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PolicyRulesWithSubjectsV1Beta1",
		reflect.TypeOf((*PolicyRulesWithSubjectsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PortworxVolumeSource",
		reflect.TypeOf((*PortworxVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Preconditions",
		reflect.TypeOf((*Preconditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PreferredSchedulingTerm",
		reflect.TypeOf((*PreferredSchedulingTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PriorityLevelConfigurationReferenceV1Beta1",
		reflect.TypeOf((*PriorityLevelConfigurationReferenceV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.PriorityLevelConfigurationSpecV1Beta1",
		reflect.TypeOf((*PriorityLevelConfigurationSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Probe",
		reflect.TypeOf((*Probe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ProjectedVolumeSource",
		reflect.TypeOf((*ProjectedVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-dns-route53.Quantity",
		reflect.TypeOf((*Quantity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Quantity{}
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.QueuingConfigurationV1Beta1",
		reflect.TypeOf((*QueuingConfigurationV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.QuobyteVolumeSource",
		reflect.TypeOf((*QuobyteVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RbdPersistentVolumeSource",
		reflect.TypeOf((*RbdPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RbdVolumeSource",
		reflect.TypeOf((*RbdVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ReplicaSetSpec",
		reflect.TypeOf((*ReplicaSetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ReplicationControllerSpec",
		reflect.TypeOf((*ReplicationControllerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ResourceAttributes",
		reflect.TypeOf((*ResourceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ResourceAttributesV1Beta1",
		reflect.TypeOf((*ResourceAttributesV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ResourceFieldSelector",
		reflect.TypeOf((*ResourceFieldSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ResourceMetricSourceV2Beta1",
		reflect.TypeOf((*ResourceMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ResourceMetricSourceV2Beta2",
		reflect.TypeOf((*ResourceMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ResourcePolicyRuleV1Beta1",
		reflect.TypeOf((*ResourcePolicyRuleV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ResourceQuotaSpec",
		reflect.TypeOf((*ResourceQuotaSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ResourceRequirements",
		reflect.TypeOf((*ResourceRequirements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RoleRef",
		reflect.TypeOf((*RoleRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RoleRefV1Alpha1",
		reflect.TypeOf((*RoleRefV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RoleRefV1Beta1",
		reflect.TypeOf((*RoleRefV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RollingUpdateDaemonSet",
		reflect.TypeOf((*RollingUpdateDaemonSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RollingUpdateDeployment",
		reflect.TypeOf((*RollingUpdateDeployment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RollingUpdateStatefulSetStrategy",
		reflect.TypeOf((*RollingUpdateStatefulSetStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RuleWithOperations",
		reflect.TypeOf((*RuleWithOperations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RuleWithOperationsV1Beta1",
		reflect.TypeOf((*RuleWithOperationsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RunAsGroupStrategyOptionsV1Beta1",
		reflect.TypeOf((*RunAsGroupStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RunAsUserStrategyOptionsV1Beta1",
		reflect.TypeOf((*RunAsUserStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RuntimeClassSpecV1Alpha1",
		reflect.TypeOf((*RuntimeClassSpecV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.RuntimeClassStrategyOptionsV1Beta1",
		reflect.TypeOf((*RuntimeClassStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ScaleIoPersistentVolumeSource",
		reflect.TypeOf((*ScaleIoPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ScaleIoVolumeSource",
		reflect.TypeOf((*ScaleIoVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ScaleSpec",
		reflect.TypeOf((*ScaleSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Scheduling",
		reflect.TypeOf((*Scheduling)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SchedulingV1Alpha1",
		reflect.TypeOf((*SchedulingV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SchedulingV1Beta1",
		reflect.TypeOf((*SchedulingV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ScopeSelector",
		reflect.TypeOf((*ScopeSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ScopedResourceSelectorRequirement",
		reflect.TypeOf((*ScopedResourceSelectorRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SeLinuxOptions",
		reflect.TypeOf((*SeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SeLinuxStrategyOptionsV1Beta1",
		reflect.TypeOf((*SeLinuxStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SeccompProfile",
		reflect.TypeOf((*SeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SecretEnvSource",
		reflect.TypeOf((*SecretEnvSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SecretKeySelector",
		reflect.TypeOf((*SecretKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SecretProjection",
		reflect.TypeOf((*SecretProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SecretReference",
		reflect.TypeOf((*SecretReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SecretVolumeSource",
		reflect.TypeOf((*SecretVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SecurityContext",
		reflect.TypeOf((*SecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SelfSubjectAccessReviewSpec",
		reflect.TypeOf((*SelfSubjectAccessReviewSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SelfSubjectAccessReviewSpecV1Beta1",
		reflect.TypeOf((*SelfSubjectAccessReviewSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SelfSubjectRulesReviewSpec",
		reflect.TypeOf((*SelfSubjectRulesReviewSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SelfSubjectRulesReviewSpecV1Beta1",
		reflect.TypeOf((*SelfSubjectRulesReviewSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ServiceAccountTokenProjection",
		reflect.TypeOf((*ServiceAccountTokenProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ServiceBackendPort",
		reflect.TypeOf((*ServiceBackendPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ServicePort",
		reflect.TypeOf((*ServicePort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ServiceReference",
		reflect.TypeOf((*ServiceReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ServiceReferenceV1Beta1",
		reflect.TypeOf((*ServiceReferenceV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ServiceSpec",
		reflect.TypeOf((*ServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SessionAffinityConfig",
		reflect.TypeOf((*SessionAffinityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.StatefulSetSpec",
		reflect.TypeOf((*StatefulSetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.StatefulSetUpdateStrategy",
		reflect.TypeOf((*StatefulSetUpdateStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.StatusCause",
		reflect.TypeOf((*StatusCause)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.StatusDetails",
		reflect.TypeOf((*StatusDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.StorageOsPersistentVolumeSource",
		reflect.TypeOf((*StorageOsPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.StorageOsVolumeSource",
		reflect.TypeOf((*StorageOsVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Subject",
		reflect.TypeOf((*Subject)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SubjectAccessReviewSpec",
		reflect.TypeOf((*SubjectAccessReviewSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SubjectAccessReviewSpecV1Beta1",
		reflect.TypeOf((*SubjectAccessReviewSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SubjectV1Alpha1",
		reflect.TypeOf((*SubjectV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SubjectV1Beta1",
		reflect.TypeOf((*SubjectV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.SupplementalGroupsStrategyOptionsV1Beta1",
		reflect.TypeOf((*SupplementalGroupsStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Sysctl",
		reflect.TypeOf((*Sysctl)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Taint",
		reflect.TypeOf((*Taint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TcpSocketAction",
		reflect.TypeOf((*TcpSocketAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TokenRequest",
		reflect.TypeOf((*TokenRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TokenRequestSpec",
		reflect.TypeOf((*TokenRequestSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TokenRequestV1Beta1",
		reflect.TypeOf((*TokenRequestV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TokenReviewSpec",
		reflect.TypeOf((*TokenReviewSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TokenReviewSpecV1Beta1",
		reflect.TypeOf((*TokenReviewSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Toleration",
		reflect.TypeOf((*Toleration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TopologySelectorLabelRequirement",
		reflect.TypeOf((*TopologySelectorLabelRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TopologySelectorTerm",
		reflect.TypeOf((*TopologySelectorTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TopologySpreadConstraint",
		reflect.TypeOf((*TopologySpreadConstraint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.TypedLocalObjectReference",
		reflect.TypeOf((*TypedLocalObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ValidatingWebhook",
		reflect.TypeOf((*ValidatingWebhook)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.ValidatingWebhookV1Beta1",
		reflect.TypeOf((*ValidatingWebhookV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.Volume",
		reflect.TypeOf((*Volume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeAttachmentSource",
		reflect.TypeOf((*VolumeAttachmentSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeAttachmentSourceV1Alpha1",
		reflect.TypeOf((*VolumeAttachmentSourceV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeAttachmentSourceV1Beta1",
		reflect.TypeOf((*VolumeAttachmentSourceV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeAttachmentSpec",
		reflect.TypeOf((*VolumeAttachmentSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeAttachmentSpecV1Alpha1",
		reflect.TypeOf((*VolumeAttachmentSpecV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeAttachmentSpecV1Beta1",
		reflect.TypeOf((*VolumeAttachmentSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeDevice",
		reflect.TypeOf((*VolumeDevice)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeMount",
		reflect.TypeOf((*VolumeMount)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeNodeAffinity",
		reflect.TypeOf((*VolumeNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeNodeResources",
		reflect.TypeOf((*VolumeNodeResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeNodeResourcesV1Beta1",
		reflect.TypeOf((*VolumeNodeResourcesV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VolumeProjection",
		reflect.TypeOf((*VolumeProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.VsphereVirtualDiskVolumeSource",
		reflect.TypeOf((*VsphereVirtualDiskVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.WebhookClientConfig",
		reflect.TypeOf((*WebhookClientConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.WebhookClientConfigV1Beta1",
		reflect.TypeOf((*WebhookClientConfigV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.WebhookConversion",
		reflect.TypeOf((*WebhookConversion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.WeightedPodAffinityTerm",
		reflect.TypeOf((*WeightedPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-dns-route53.WindowsSecurityContextOptions",
		reflect.TypeOf((*WindowsSecurityContextOptions)(nil)).Elem(),
	)
}
