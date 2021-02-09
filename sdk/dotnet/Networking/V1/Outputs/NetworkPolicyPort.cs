// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Networking.V1
{

    [OutputType]
    public sealed class NetworkPolicyPort
    {
        /// <summary>
        /// If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port. This feature is in Alpha state and should be enabled using the Feature Gate "NetworkPolicyEndPort".
        /// </summary>
        public readonly int EndPort;
        /// <summary>
        /// The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
        /// </summary>
        public readonly Union<int, string> Port;
        /// <summary>
        /// The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private NetworkPolicyPort(
            int endPort,

            Union<int, string> port,

            string protocol)
        {
            EndPort = endPort;
            Port = port;
            Protocol = protocol;
        }
    }
}
