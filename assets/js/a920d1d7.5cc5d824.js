"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9369],{1233:(e,n,i)=>{i.r(n),i.d(n,{assets:()=>l,contentTitle:()=>d,default:()=>h,frontMatter:()=>r,metadata:()=>s,toc:()=>o});const s=JSON.parse('{"id":"built-in-services/data-svc","title":"Data Svc","description":"The Data Service (Data Svc) is designed to facilitate backendless applications, allowing data to be saved and queried directly from the frontend, similar to Firebase.","source":"@site/docs/built-in-services/data-svc.md","sourceDirName":"built-in-services","slug":"/built-in-services/data-svc","permalink":"/docs/built-in-services/data-svc","draft":false,"unlisted":false,"editUrl":"https://github.com/1backend/1backend/tree/main/docs-source/docs/built-in-services/data-svc.md","tags":[{"inline":true,"label":"data-svc","permalink":"/docs/tags/data-svc"},{"inline":true,"label":"data","permalink":"/docs/tags/data"},{"inline":true,"label":"permissions","permalink":"/docs/tags/permissions"},{"inline":true,"label":"roles","permalink":"/docs/tags/roles"},{"inline":true,"label":"authentication","permalink":"/docs/tags/authentication"},{"inline":true,"label":"authorization","permalink":"/docs/tags/authorization"},{"inline":true,"label":"services","permalink":"/docs/tags/services"}],"version":"current","sidebarPosition":100,"frontMatter":{"sidebar_position":100,"tags":["data-svc","data","permissions","roles","authentication","authorization","services"]},"sidebar":"tutorialSidebar","previous":{"title":"Model Svc","permalink":"/docs/built-in-services/model-svc"},"next":{"title":"Container Svc","permalink":"/docs/built-in-services/container-svc"}}');var t=i(74848),a=i(28453);const r={sidebar_position:100,tags:["data-svc","data","permissions","roles","authentication","authorization","services"]},d="Data Svc",l={},o=[{value:"Purpose",id:"purpose",level:2},{value:"Data types",id:"data-types",level:2},{value:"Objects",id:"objects",level:2},{value:"Data model",id:"data-model",level:3},{value:"Permission model",id:"permission-model",level:3},{value:"Readers",id:"readers",level:2},{value:"Authors",id:"authors",level:2},{value:"Writers",id:"writers",level:2},{value:"Deleters",id:"deleters",level:2},{value:"Conventions",id:"conventions",level:3},{value:"Table name and object ID",id:"table-name-and-object-id",level:4},{value:"<code>_self</code>",id:"_self",level:4},{value:"Querying",id:"querying",level:2},{value:"Ordering by descending value",id:"ordering-by-descending-value",level:3},{value:"Ordering by ascending value",id:"ordering-by-ascending-value",level:3},{value:"Limiting results",id:"limiting-results",level:3},{value:"Paginating with after",id:"paginating-with-after",level:3},{value:"Paginating with after in descending order",id:"paginating-with-after-in-descending-order",level:3}];function c(e){const n={a:"a",blockquote:"blockquote",code:"code",h1:"h1",h2:"h2",h3:"h3",h4:"h4",header:"header",li:"li",p:"p",pre:"pre",ul:"ul",...(0,a.R)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(n.header,{children:(0,t.jsx)(n.h1,{id:"data-svc",children:"Data Svc"})}),"\n",(0,t.jsx)(n.p,{children:"The Data Service (Data Svc) is designed to facilitate backendless applications, allowing data to be saved and queried directly from the frontend, similar to Firebase."}),"\n",(0,t.jsxs)(n.blockquote,{children:["\n",(0,t.jsxs)(n.p,{children:["This page provides a high-level overview of ",(0,t.jsx)(n.code,{children:"Data Svc"}),". For detailed information, refer to the ",(0,t.jsx)(n.a,{href:"/docs/1backend/query-objects",children:"Data Svc API documentation"}),"."]}),"\n"]}),"\n",(0,t.jsx)(n.h2,{id:"purpose",children:"Purpose"}),"\n",(0,t.jsx)(n.p,{children:"Data Svc serves as a lightweight database abstraction designed for rapid prototyping. It allows direct reading, writing, and deletion from the frontend, eliminating the need for basic CRUD microservices that merely handle routine database operations."}),"\n",(0,t.jsx)(n.h2,{id:"data-types",children:"Data types"}),"\n",(0,t.jsxs)(n.p,{children:["Currently, Data Svc supports only untyped, dynamic, and schemaless entries known as ",(0,t.jsx)(n.code,{children:"Object"}),"s."]}),"\n",(0,t.jsx)(n.h2,{id:"objects",children:"Objects"}),"\n",(0,t.jsx)(n.h3,{id:"data-model",children:"Data model"}),"\n",(0,t.jsx)(n.p,{children:"Multiple tenants (users or services) write to shared tables. Access is governed by the permission model outlined below."}),"\n",(0,t.jsx)(n.h3,{id:"permission-model",children:"Permission model"}),"\n",(0,t.jsxs)(n.p,{children:["The ",(0,t.jsx)(n.code,{children:"Data Svc"})," ",(0,t.jsx)(n.code,{children:"Object"})," permission model is designed with two primary goals:"]}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"Simplicity \u2013 Easy to understand and implement"}),"\n",(0,t.jsx)(n.li,{children:"Flexibility \u2013 Versatile while maintaining simplicity"}),"\n"]}),"\n",(0,t.jsx)(n.p,{children:"To illustrate the model, consider the following example entry:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-yaml",children:'table: "pet"\nid: "pet_67890"\ndata:\n  yourCustomKey1: "yourCustomValue1"\n  yourCustomKey2: 42\nreaders:\n  - "usr_12345"\n  - "org_67890"\nwriters:\n  - "org_67890"\ndeleters:\n  - "usr_12345"\nauthors:\n  - "usr_99999"\n  - "org_99999"\n'})}),"\n",(0,t.jsx)(n.h2,{id:"readers",children:"Readers"}),"\n",(0,t.jsxs)(n.p,{children:["The ",(0,t.jsx)(n.code,{children:"readers"})," field defines which users, organizations, or roles have permission to view an entry."]}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"Users and organizations outside of your own can be permited access."}),"\n",(0,t.jsx)(n.li,{children:"This field can be set by the author or writers to include user IDs, organization IDs, or roles they themselves do not belong to."}),"\n"]}),"\n",(0,t.jsx)(n.h2,{id:"authors",children:"Authors"}),"\n",(0,t.jsxs)(n.p,{children:["The ",(0,t.jsx)(n.code,{children:"authors"})," field identifies the original creators of an entry. Unlike ",(0,t.jsx)(n.code,{children:"readers"}),", ",(0,t.jsx)(n.code,{children:"writers"}),", and ",(0,t.jsx)(n.code,{children:"deleters"}),", which are user-defined, this field is system-assigned. It can only include the author's user ID, organization IDs they belong to, or roles they hold. This ensures it cannot be altered or spoofed, helping to prevent spam."]}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:'In multi-tenant applications, spam can occur because anyone can "offer" a record to be read by another user or organization they are not part of. This can be problematic\u2014for example, in a chat application where strangers could send unsolicited messages simply by knowing a company ID.'}),"\n",(0,t.jsxs)(n.li,{children:["The ",(0,t.jsx)(n.code,{children:"authors"})," field helps prevent such abuse limiting the list to resources the author has."]}),"\n"]}),"\n",(0,t.jsx)(n.h2,{id:"writers",children:"Writers"}),"\n",(0,t.jsxs)(n.p,{children:["The ",(0,t.jsx)(n.code,{children:"writers"})," field specifies which users, organizations, or roles have permission to modify an entry."]}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"This field can be set by the author or existing writers to include user IDs, organization IDs, or roles they themselves do not belong to."}),"\n"]}),"\n",(0,t.jsx)(n.h2,{id:"deleters",children:"Deleters"}),"\n",(0,t.jsxs)(n.p,{children:["The ",(0,t.jsx)(n.code,{children:"deleters"})," field defines which users, organizations, or roles are authorized to delete an entry."]}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"This field can be set by the author or existing writers to include user IDs, organization IDs, or roles they themselves do not belong to."}),"\n"]}),"\n",(0,t.jsx)(n.h3,{id:"conventions",children:"Conventions"}),"\n",(0,t.jsx)(n.h4,{id:"table-name-and-object-id",children:"Table name and object ID"}),"\n",(0,t.jsx)(n.p,{children:"Each object ID must be prefixed with the table name. Whenever possible, use singular table names."}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-yaml",children:'table: "pet"\nid: "pet_67890"\n'})}),"\n",(0,t.jsx)(n.h4,{id:"_self",children:(0,t.jsx)(n.code,{children:"_self"})}),"\n",(0,t.jsxs)(n.p,{children:["You can specify the reserved string ",(0,t.jsx)(n.code,{children:"_self"})," in the ",(0,t.jsx)(n.code,{children:"readers"}),", ",(0,t.jsx)(n.code,{children:"writers"})," or ",(0,t.jsx)(n.code,{children:"deleters"})," lists. It will be extrapolated to your user ID."]}),"\n",(0,t.jsx)(n.h2,{id:"querying",children:"Querying"}),"\n",(0,t.jsx)(n.p,{children:"Data Svc allows querying objects with flexible filtering, sorting, and pagination options."}),"\n",(0,t.jsx)(n.h3,{id:"ordering-by-descending-value",children:"Ordering by descending value"}),"\n",(0,t.jsx)(n.p,{children:"Request"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "table": "pet",\n  "query": {\n    "orderBys": [\n      {\n        "field": "age",\n        "desc": true,\n        "sortingType": "numeric"\n      }\n    ]\n  }\n}\n'})}),"\n",(0,t.jsxs)(n.p,{children:["You might wonder what sorting type is. It is essentially a clutch for some systems like PostgreSQL, where the Data Svc and its ORM stores the dynamic fields of Objects in a ",(0,t.jsx)(n.code,{children:"JSONB"})," field."]}),"\n",(0,t.jsxs)(n.p,{children:["Unfortunately ordering on JSONB fields defaults to string sorting. The ",(0,t.jsx)(n.code,{children:"sortingType"})," field helps the system force the correct ordering. For possible ordering values, see the ",(0,t.jsx)(n.a,{href:"/docs/1backend/query-objects",children:"queryObjects endpoint API"})]}),"\n",(0,t.jsx)(n.p,{children:"Response:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "objects": [\n    { "table": "pet", "id": "pet_19", "data": { "age": 19 } },\n    { "table": "pet", "id": "pet_18", "data": { "age": 18 } },\n    { "table": "pet", "id": "pet_17", "data": { "age": 17 } }\n  ]\n}\n'})}),"\n",(0,t.jsx)(n.h3,{id:"ordering-by-ascending-value",children:"Ordering by ascending value"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "table": "pet",\n  "query": {\n    "orderBys": [\n      {\n        "field": "age",\n        "sortingType": "numeric"\n      }\n    ]\n  }\n}\n'})}),"\n",(0,t.jsx)(n.p,{children:"Response:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "objects": [\n    { "table": "pet", "id": "pet_0", "data": { "age": 0 } },\n    { "table": "pet", "id": "pet_1", "data": { "age": 1 } },\n    { "table": "pet", "id": "pet_2", "data": { "age": 2 } }\n  ]\n}\n'})}),"\n",(0,t.jsx)(n.h3,{id:"limiting-results",children:"Limiting results"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "table": "pet",\n  "query": {\n    "orderBys": [\n      {\n        "field": "age",\n        "sortingType": "numeric"\n      }\n    ],\n    "limit": 5\n  }\n}\n'})}),"\n",(0,t.jsx)(n.p,{children:"Response:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "objects": [\n    { "table": "pet", "id": "pet_0", "data": { "age": 0 } },\n    { "table": "pet", "id": "pet_1", "data": { "age": 1 } },\n    { "table": "pet", "id": "pet_2", "data": { "age": 2 } },\n    { "table": "pet", "id": "pet_3", "data": { "age": 3 } },\n    { "table": "pet", "id": "pet_4", "data": { "age": 4 } }\n  ]\n}\n'})}),"\n",(0,t.jsx)(n.h3,{id:"paginating-with-after",children:"Paginating with after"}),"\n",(0,t.jsx)(n.p,{children:"Request:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "table": "pet",\n  "query": {\n    "orderBys": [\n      {\n        "field": "age",\n        "sortingType": "numeric"\n      }\n    ],\n    "limit": 5,\n    "jsonAfter": "[4]"\n  }\n}\n'})}),"\n",(0,t.jsxs)(n.p,{children:["The ",(0,t.jsx)(n.code,{children:"after"})," field is named ",(0,t.jsx)(n.code,{children:"jsonAfter"})," and is a string-marshalled array to address limitations in Swaggo. Specifically, Swaggo converts []interface"," to []map[string]interface",", and there is no way to define a type that resolves to an any/interface"," during the ",(0,t.jsx)(n.code,{children:"go -> openapi -> go"})," generation process. Therefore, ",(0,t.jsx)(n.code,{children:"jsonAfter"})," is a JSON-encoded string representing an array, e.g., ",(0,t.jsx)(n.code,{children:"[42]"}),"."]}),"\n",(0,t.jsx)(n.p,{children:"Response:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "objects": [\n    { "table:" "pet", "id": "pet_5", "data": { "age": 5 } },\n    { "table:" "pet", "id": "pet_6", "data": { "age": 6 } },\n    { "table:" "pet", "id": "pet_7", "data": { "age": 7 } },\n    { "table:" "pet", "id": "pet_8", "data": { "age": 8 } },\n    { "table:" "pet", "id": "pet_9", "data": { "age": 9 } }\n  ]\n}\n'})}),"\n",(0,t.jsx)(n.h3,{id:"paginating-with-after-in-descending-order",children:"Paginating with after in descending order"}),"\n",(0,t.jsx)(n.p,{children:"Request:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "table": "pet",\n  "query": {\n    "orderBys": [\n      {\n        "field": "age",\n        "desc": true,\n        "sortingType": "numeric"\n      }\n    ],\n    "limit": 5,\n    "jsonAfter": "[15]"\n  }\n}\n'})}),"\n",(0,t.jsx)(n.p,{children:"Response:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'{\n  "objects": [\n    { "id": "pet_14", "data": { "age": 14 } },\n    { "id": "pet_13", "data": { "age": 13 } },\n    { "id": "pet_12", "data": { "age": 12 } },\n    { "id": "pet_11", "data": { "age": 11 } },\n    { "id": "pet_10", "data": { "age": 10 } }\n  ]\n}\n'})})]})}function h(e={}){const{wrapper:n}={...(0,a.R)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(c,{...e})}):c(e)}},28453:(e,n,i)=>{i.d(n,{R:()=>r,x:()=>d});var s=i(96540);const t={},a=s.createContext(t);function r(e){const n=s.useContext(a);return s.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function d(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:r(e.components),s.createElement(a.Provider,{value:n},e.children)}}}]);