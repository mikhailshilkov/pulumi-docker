// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages the configuration of a Docker service in a swarm.
 * 
 * ## Basic
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 * 
 * // Creates a config
 * const fooConfig = new docker.Config("fooConfig", {
 *     data: "ewogICJzZXJIfQo=",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/r/config.html.markdown.
 */
export class Config extends pulumi.CustomResource {
    /**
     * Get an existing Config resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigState, opts?: pulumi.CustomResourceOptions): Config {
        return new Config(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:index/config:Config';

    /**
     * Returns true if the given object is an instance of Config.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Config {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Config.__pulumiType;
    }

    /**
     * The base64 encoded data of the config.
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * The name of the Docker config.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Config resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigArgs | ConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConfigState | undefined;
            inputs["data"] = state ? state.data : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ConfigArgs | undefined;
            if (!args || args.data === undefined) {
                throw new Error("Missing required property 'data'");
            }
            inputs["data"] = args ? args.data : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Config.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Config resources.
 */
export interface ConfigState {
    /**
     * The base64 encoded data of the config.
     */
    readonly data?: pulumi.Input<string>;
    /**
     * The name of the Docker config.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Config resource.
 */
export interface ConfigArgs {
    /**
     * The base64 encoded data of the config.
     */
    readonly data: pulumi.Input<string>;
    /**
     * The name of the Docker config.
     */
    readonly name?: pulumi.Input<string>;
}
