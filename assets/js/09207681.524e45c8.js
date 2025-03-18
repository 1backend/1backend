"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4770],{71131:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>m,contentTitle:()=>g,default:()=>v,frontMatter:()=>j,metadata:()=>r,toc:()=>y});const r=JSON.parse('{"id":"1backend/update-objects","title":"Update Objects","description":"Update fields of objects that match the given filters using the provided object.","source":"@site/docs/1backend/update-objects.api.mdx","sourceDirName":"1backend","slug":"/1backend/update-objects","permalink":"/docs/1backend/update-objects","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"update-objects","title":"Update Objects","description":"Update fields of objects that match the given filters using the provided object.","sidebar_label":"Update Objects","hide_title":true,"hide_table_of_contents":true,"api":"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","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Delete Objects","permalink":"/docs/1backend/delete-objects"},"next":{"title":"Upsert Objects","permalink":"/docs/1backend/upsert-objects"}}');var n=s(74848),o=s(28453),a=s(53746),i=s.n(a),d=s(56518),c=s.n(d),p=s(99972),l=s.n(p),u=s(25342),b=s.n(u),h=(s(44215),s(82223),s(24861));const j={id:"update-objects",title:"Update Objects",description:"Update fields of objects that match the given filters using the provided object.",sidebar_label:"Update Objects",hide_title:!0,hide_table_of_contents:!0,api:"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",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},g=void 0,m={},y=[];function D(e){const t={p:"p",...(0,o.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Update Objects"}),"\n",(0,n.jsx)(i(),{method:"post",path:"/data-svc/objects/update",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Update fields of objects that match the given filters using the provided object.\nAny fields not included in the incoming object will remain unchanged."}),"\n",(0,n.jsx)(h.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(c(),{parameters:void 0}),"\n",(0,n.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{filters:{description:"Filters to determine which objects will be updated.\nOnly objects matching all filters will be modified.",items:{properties:{fields:{items:{type:"string"},type:"array"},jsonValues:{description:"JSONValues is a JSON marshalled array of values.\nIt's JSON marhalled due to the limitations of the\nSwaggo -> OpenAPI 2.0 -> OpenAPI Go generator toolchain.",type:"string"},op:{enum:["equals","containsSubstring","startsWith","intersects","isInList"],type:"string","x-enum-varnames":["OpEquals","OpContainsSubstring","OpStartsWith","OpIntersects","OpIsInList"]}},type:"object"},type:"array"},object:{allOf:[{properties:{authors:{description:"Authors is a list of user ID and organization ID who created the object.\nThe authors field tracks which users or organizations created an entry, helping to prevent spam.\nIf an organization ID is not provided, the currently active organization will\nbe queried from the User Svc.",example:['["usr_12345"',' "org_67890"]'],items:{type:"string"},type:"array"},createdAt:{type:"string"},data:{additionalProperties:!0,type:"object"},deleters:{description:"Deleters is a list of user IDs and role IDs that can delete the object.\n`_self` can be used to refer to the caller user's userId and\n`_org` can be used to refer to the user's currently active organization (if exists).",example:['["usr_12345"',' "org_67890"]'],items:{type:"string"},type:"array"},id:{type:"string"},readers:{description:"Readers is a list of user IDs and role IDs that can read the object.\n`_self` can be used to refer to the caller user's userId and\n`_org` can be used to refer to the user's currently active organization (if exists).",example:['["usr_12345"',' "org_67890"]'],items:{type:"string"},type:"array"},table:{type:"string"},updatedAt:{type:"string"},writers:{description:"Writers is a list of user IDs and role IDs that can write the object.\n`_self` can be used to refer to the caller user's userId and\n`_org` can be used to refer to the user's currently active organization (if exists).",example:['["usr_12345"',' "org_67890"]'],items:{type:"string"},type:"array"}},required:["data","table"],type:"object"}],description:"The object containing the fields to update in matching objects."},table:{type:"string"}},type:"object"}}},description:"Update request payload",required:!0}}),"\n",(0,n.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"Successful update of objects",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function v(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(D,{...e})}):D(e)}}}]);