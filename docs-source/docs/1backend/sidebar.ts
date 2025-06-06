import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebar: SidebarsConfig = {
  apisidebar: [
    {
      type: "doc",
      id: "1backend/1-backend",
    },
    {
      type: "category",
      label: "Chat Svc",
      items: [
        {
          type: "doc",
          id: "1backend/events",
          label: "Events",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/delete-message",
          label: "Delete a Message",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/list-messages",
          label: "List Messages",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/save-thread",
          label: "Save Thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/delete-thread",
          label: "Delete a Thread",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/save-message",
          label: "Save Message",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/list-threads",
          label: "List Threads",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Config Svc",
      items: [
        {
          type: "doc",
          id: "1backend/read-config",
          label: "Read Config",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/save-config",
          label: "Save Config",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "Container Svc",
      items: [
        {
          type: "doc",
          id: "1backend/run-container",
          label: "Run a Container",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/container-is-running",
          label: "Check If a Container Is Running",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/stop-container",
          label: "Stop a Container",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/container-summary",
          label: "Get Container Summary",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/list-containers",
          label: "List Containers",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/container-daemon-info",
          label: "Get Container Daemon Information",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/get-host",
          label: "Get Container Host",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/build-image",
          label: "Build an Image",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/image-pullable",
          label: "Check if Container Image is Pullable",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/list-container-logs",
          label: "List Logs",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Data Svc",
      items: [
        {
          type: "doc",
          id: "1backend/create-object",
          label: "Create a Generic Object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/upsert-object",
          label: "Upsert a Generic Object",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/query-objects",
          label: "Query Objects",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/delete-objects",
          label: "Delete Objects",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/update-objects",
          label: "Update Objects",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/upsert-objects",
          label: "Upsert Objects",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "Deploy Svc",
      items: [
        {
          type: "doc",
          id: "1backend/delete-deployment",
          label: "Delete Deployment",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/save-deployment",
          label: "Save Deployment",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/list-deployments",
          label: "List Deployments",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Email Svc",
      items: [
        {
          type: "doc",
          id: "1backend/send-email",
          label: "Send an Email",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "File Svc",
      items: [
        {
          type: "doc",
          id: "1backend/download-file",
          label: "Download a File",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/get-download",
          label: "Get a Download",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/pause-download",
          label: "Pause a Download",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/list-downloads",
          label: "List Downloads",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/serve-download",
          label: "Serve a Downloaded file",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/serve-upload",
          label: "Serve an Uploaded File",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/upload-file",
          label: "Upload a File",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/list-uploads",
          label: "List Uploads",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Firehose Svc",
      items: [
        {
          type: "doc",
          id: "1backend/publish-event",
          label: "Publish an Event",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/subscribe-to-events",
          label: "Subscribe to the Event Stream",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "Image Svc",
      items: [
        {
          type: "doc",
          id: "1backend/serve-uploaded-image",
          label: "Serve Uploaded Image",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "Model Svc",
      items: [
        {
          type: "doc",
          id: "1backend/start-default-model",
          label: "Start the Default Model",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/get-default-model-status",
          label: "Get Default Model Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/get-model",
          label: "Get a Model",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/make-default",
          label: "Make a Model Default",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/start-model",
          label: "Start a Model",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/get-model-status",
          label: "Get Model Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/list-models",
          label: "List Models",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/list-platforms",
          label: "List Platforms",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Policy Svc",
      items: [
        {
          type: "doc",
          id: "1backend/check",
          label: "Check",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/upsert-instance",
          label: "Upsert an Instance",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "Prompt Svc",
      items: [
        {
          type: "doc",
          id: "1backend/prompt",
          label: "Prompt an AI",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/list-prompts",
          label: "List Prompts",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/subscribe-to-prompt-responses",
          label: "Subscribe to Prompt Responses by Thread",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/remove-prompt",
          label: "Remove Prompt",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/prompt-types",
          label: "Prompt Types",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Proxy Svc",
      items: [
        {
          type: "doc",
          id: "1backend/list-certs",
          label: "List Certs",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/list-routes",
          label: "List Routes",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/save-routes",
          label: "Save Routes",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "Registry Svc",
      items: [
        {
          type: "doc",
          id: "1backend/save-definition",
          label: "Register a Definition",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/delete-definition",
          label: "Delete Definition",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/list-definitions",
          label: "List Definitions",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/register-instance",
          label: "Register Instance",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/remove-instance",
          label: "Remove Instance",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/list-instances",
          label: "List Service Instances",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/delete-node",
          label: "Delete Node",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/self-node",
          label: "View Self Node",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/list-nodes",
          label: "List Nodes",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Secret Svc",
      items: [
        {
          type: "doc",
          id: "1backend/decrypt-value",
          label: "Decrypt a Value",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/encrypt-value",
          label: "Encrypt a Value",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/is-secure",
          label: "Check Security Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/remove-secrets",
          label: "Remove Secrets",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/list-secrets",
          label: "List Secrets",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/save-secrets",
          label: "Save Secrets",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "Source Svc",
      items: [
        {
          type: "doc",
          id: "1backend/checkout-repo",
          label: "Checkout a git repository",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "User Svc",
      items: [
        {
          type: "doc",
          id: "1backend/reset-password",
          label: "Reset Password",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/change-password",
          label: "Change Password",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/list-enrolls",
          label: "List Enrolls",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/save-enrolls",
          label: "Save Enrolls",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/login",
          label: "Login",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/save-organization",
          label: "Save an Organization",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/delete-membership",
          label: "Delete Membership",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/save-membership",
          label: "Save Membership",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/list-organizations",
          label: "List Organizations",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/list-permissions",
          label: "List Permissions",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/list-permits",
          label: "List Permits",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/save-permits",
          label: "Save Permits",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/get-public-key",
          label: "Get Public Key",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "1backend/refresh-token",
          label: "Refresh Token",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/register",
          label: "Register",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/read-self",
          label: "Read Self",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/save-self",
          label: "Save User Profile",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/has-permission",
          label: "Has Permission",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/revoke-tokens",
          label: "Revoke Tokens",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/create-user",
          label: "Create a New User",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "1backend/delete-user",
          label: "Delete a User",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "1backend/save-user",
          label: "Save User",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "1backend/list-users",
          label: "List Users",
          className: "api-method post",
        },
      ],
    },
  ],
};

export default sidebar.apisidebar;
