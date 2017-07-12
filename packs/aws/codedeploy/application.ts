// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

export class Application extends lumi.NamedResource implements ApplicationArgs {
    public readonly _name: string;
    public readonly uniqueId?: string;

    constructor(name: string, args: ApplicationArgs) {
        super(name);
        this._name = args._name;
        this.uniqueId = args.uniqueId;
    }
}

export interface ApplicationArgs {
    readonly _name: string;
    readonly uniqueId?: string;
}
