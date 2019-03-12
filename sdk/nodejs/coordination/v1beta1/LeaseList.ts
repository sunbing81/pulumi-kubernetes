// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";

    /**
     * LeaseList is a list of Lease objects.
     */
    export class LeaseList extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"coordination.k8s.io/v1beta1">;

      /**
       * Items is a list of schema objects.
       */
      public readonly items: pulumi.Output<outputApi.coordination.v1beta1.Lease[]>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"LeaseList">;

      /**
       * Standard list metadata. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ListMeta>;

      /**
       * Get the state of an existing `LeaseList` resource, as identified by `id`.
       * Typically this ID  is of the form <namespace>/<name>; if <namespace> is omitted, then (per
       * Kubernetes convention) the ID becomes default/<name>.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form
       *  <namespace>/<name> or <name>.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LeaseList {
          return new LeaseList(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.coordination.v1beta1.LeaseList { return this.__inputs; }
      private readonly __inputs: inputApi.coordination.v1beta1.LeaseList;

      /**
       * Create a coordination.v1beta1.LeaseList resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.coordination.v1beta1.LeaseList, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "coordination.k8s.io/v1beta1";
          inputs["items"] = args && args.items || undefined;
          inputs["kind"] = "LeaseList";
          inputs["metadata"] = args && args.metadata || undefined;
          super("kubernetes:coordination.k8s.io/v1beta1:LeaseList", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }
