/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class FileSvcUpload {
    'createdAt'?: string;
    /**
    * Logical file ID spanning all replicas
    */
    'fileId'?: string;
    /**
    * Filename is the original name of the file
    */
    'fileName'?: string;
    /**
    * FilePath is the full node local path of the file
    */
    'filePath'?: string;
    'fileSize': number;
    /**
    * Unique ID for this replica
    */
    'id'?: string;
    /**
    * ID of the node storing this replica
    */
    'nodeId'?: string;
    'updatedAt'?: string;
    'userId'?: string;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
