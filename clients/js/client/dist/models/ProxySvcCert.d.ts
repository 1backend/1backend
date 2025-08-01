/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ProxySvcCert
 */
export interface ProxySvcCert {
    /**
     * PEM-encoded certificate bundle
     *
     *  -----BEGIN EC PARAMETERS-----
     *  BggqhkjOPQMBBw==
     *  -----END EC PARAMETERS-----
     *  -----BEGIN EC PRIVATE KEY-----
     *  MHcCAQEEIDC3+7pySTQl6WRBuef...
     *  -----END EC PRIVATE KEY-----
     *  -----BEGIN CERTIFICATE-----
     *  MIIBhTCCASugAwIBAgIUQYwE...
     *  -----END CERTIFICATE-----
     * @type {string}
     * @memberof ProxySvcCert
     */
    cert: string;
    /**
     * Subject Common Name (typically domain)
     * @type {string}
     * @memberof ProxySvcCert
     */
    commonName?: string;
    /**
     * When cert record was created
     * @type {string}
     * @memberof ProxySvcCert
     */
    createdAt: string;
    /**
     * Subject Alternative Names (covered domains)
     * @type {Array<string>}
     * @memberof ProxySvcCert
     */
    dnsNames?: Array<string>;
    /**
     * Id is the host which this cert is for, e.g., "example.com" or "www.example.com"
     * @type {string}
     * @memberof ProxySvcCert
     */
    id: string;
    /**
     * Whether cert is a CA (usually false for LE certs)
     * @type {boolean}
     * @memberof ProxySvcCert
     */
    isCA?: boolean;
    /**
     * Certificate issuer name (e.g., Let's Encrypt)
     * @type {string}
     * @memberof ProxySvcCert
     */
    issuer?: string;
    /**
     * Cert validity end time
     * @type {string}
     * @memberof ProxySvcCert
     */
    notAfter?: string;
    /**
     * Cert validity start time
     * @type {string}
     * @memberof ProxySvcCert
     */
    notBefore?: string;
    /**
     * Public key algorithm (e.g., RSA, ECDSA)
     * @type {string}
     * @memberof ProxySvcCert
     */
    publicKeyAlgorithm?: string;
    /**
     * Bit length of the public key
     * @type {number}
     * @memberof ProxySvcCert
     */
    publicKeyBitLength?: number;
    /**
     * Unique certificate serial number
     * @type {string}
     * @memberof ProxySvcCert
     */
    serialNumber?: string;
    /**
     * Algorithm used to sign the cert (e.g., SHA256-RSA)
     * @type {string}
     * @memberof ProxySvcCert
     */
    signatureAlgorithm?: string;
    /**
     * When cert record was last updated
     * @type {string}
     * @memberof ProxySvcCert
     */
    updatedAt: string;
}
/**
 * Check if a given object implements the ProxySvcCert interface.
 */
export declare function instanceOfProxySvcCert(value: object): value is ProxySvcCert;
export declare function ProxySvcCertFromJSON(json: any): ProxySvcCert;
export declare function ProxySvcCertFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProxySvcCert;
export declare function ProxySvcCertToJSON(json: any): ProxySvcCert;
export declare function ProxySvcCertToJSONTyped(value?: ProxySvcCert | null, ignoreDiscriminator?: boolean): any;
