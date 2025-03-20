"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9173],{6325:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>j,contentTitle:()=>b,default:()=>v,frontMatter:()=>g,metadata:()=>s,toc:()=>y});const s=JSON.parse('{"id":"1backend/create-role","title":"Create a New Role","description":"Create a new role.","source":"@site/docs/1backend/create-role.api.mdx","sourceDirName":"1backend","slug":"/1backend/create-role","permalink":"/docs/1backend/create-role","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"create-role","title":"Create a New Role","description":"Create a new role.","sidebar_label":"Create a New Role","hide_title":true,"hide_table_of_contents":true,"api":"eJztVk1v4zYQ/SsEL70okpN0gVZYLJq0QeF2sQni5BQbCE2NJW4okiEpK15B/70YUrKd2FsUi3ZPvRgWOR9v5nEe2dECHLfCeKEVzemvFpgHwoiCllgtIZ2r98sPdxWELzL9jdSN82QJxFhYiRcoyHJDfAWEMynB/uCIk02Zvs+WH+bqqkyJWB1uE+HIowHvvLZw4tb8MdiEFK7SjSyI1PqJSPEErw1zVtRCPaZzhZgaB5YwznWjPGkrTXjA73bRWiElWQLXNYRF3SqwRCMm5oNJQpgqiFZys2cwuBG2lEC8JlAIvw2aztVc3cJzI+yQ6hGBBHi4n0cUj8SArYVzQquUJlQbsAz7PC1oTqPNrZZAE2rhuQHnL3WxoXlHuVYelMe/zBgpeHDLPjvkqKOOV1Az/GcsBvUCHH69orKjfmOA5tR5K1RJ+4SK4uiyYjUc3djBnxYhgfBQu6OmwwKzlm1o38eKhIWC5g+Yd7E10cvPwD3t0ejo2cOWkNvYELofyNsGQmRntHKx5LPJ5KByGiLE/hbENZyDc6tGyg1NvrW1SOvh6pDjwh/tyb9NRziZ0+NOjSm+iqQ/7P0RNvqE/nismVO1ZlIU5I/Z9adv7x9Yq+0/wxaRnB4iuVes8ZW24gsU3w3Ju+M98WAVk2QGdg2WXIWY3wdSn1AHvLHCb2j+0NFLYBbsReMrmj8sehw0Vjocu3sUx9ma4/DV4CuNsmN0mCrD0J5mo25lNgqRCwW5ELmxEk0yqTmTlXY+f/fT2fkpxRwjhBmWFKvYB/K2YXcbA2Q+mMwpWWkpdRvvDkacYRyCCnv9BIowHkeerKyug76OpZCPuhSKgCqMFsqjqgqMXwErABmIw0MvhoMS+k938mTEnxD1CTm53anu1QurTRzx18gHLuK07r6GPNvvN1L5MO4scM7VSo+qzng4HVAzgc11TIL7xQlVNpJ5q1XKdb1Xxs2UzBpjtEXOIh+V9ybPstMl40+gCnTIDtSGXkxPFPNiDaQW3GqkVXBwxEjmV9rW2DkpOCgXah7z/X7zkazP08mrbC7PsrZt01I1qbZlNvi5jJVGnpynk7TytQy3ANjaXa9mMdsOrGtZWYJNhc6CSYaUCI/9pqeXsRCaUDx5Ef4kPU8nJ5anZz+Hi0g7XzO1h3T7TvkELRku0Td6ux3G/x81/+GjZpgsDy8+M5IJFW4jPDvdIDIPdIyDtzlStUgoygludd2SObi3su9x+bkBi7K2SOiaWYFAg6ol44SjLj3BBkmN9J6gtODZYbKJI/5Gd/tk9LjgHIz/W9t9pby5nt3RhC6HN1mtC/SxrMUyWEtzGt506B2fCLjWUclU2bASbWNMVBs2qOJOhxBSMv7BqsYttdlD+FbHYiH4i2Uddem6qHJ9v7WPW1/12IpntEYKF33f/wX7ETKa","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Register","permalink":"/docs/1backend/register"},"next":{"title":"Delete a Role","permalink":"/docs/1backend/delete-role"}}');var n=r(74848),i=r(28453),o=r(53746),a=r.n(o),l=r(56518),c=r.n(l),d=r(99972),p=r.n(d),u=r(25342),h=r.n(u),m=(r(44215),r(82223),r(24861));const g={id:"create-role",title:"Create a New Role",description:"Create a new role.",sidebar_label:"Create a New Role",hide_title:!0,hide_table_of_contents:!0,api:"eJztVk1v4zYQ/SsEL70okpN0gVZYLJq0QeF2sQni5BQbCE2NJW4okiEpK15B/70YUrKd2FsUi3ZPvRgWOR9v5nEe2dECHLfCeKEVzemvFpgHwoiCllgtIZ2r98sPdxWELzL9jdSN82QJxFhYiRcoyHJDfAWEMynB/uCIk02Zvs+WH+bqqkyJWB1uE+HIowHvvLZw4tb8MdiEFK7SjSyI1PqJSPEErw1zVtRCPaZzhZgaB5YwznWjPGkrTXjA73bRWiElWQLXNYRF3SqwRCMm5oNJQpgqiFZys2cwuBG2lEC8JlAIvw2aztVc3cJzI+yQ6hGBBHi4n0cUj8SArYVzQquUJlQbsAz7PC1oTqPNrZZAE2rhuQHnL3WxoXlHuVYelMe/zBgpeHDLPjvkqKOOV1Az/GcsBvUCHH69orKjfmOA5tR5K1RJ+4SK4uiyYjUc3djBnxYhgfBQu6OmwwKzlm1o38eKhIWC5g+Yd7E10cvPwD3t0ejo2cOWkNvYELofyNsGQmRntHKx5LPJ5KByGiLE/hbENZyDc6tGyg1NvrW1SOvh6pDjwh/tyb9NRziZ0+NOjSm+iqQ/7P0RNvqE/nismVO1ZlIU5I/Z9adv7x9Yq+0/wxaRnB4iuVes8ZW24gsU3w3Ju+M98WAVk2QGdg2WXIWY3wdSn1AHvLHCb2j+0NFLYBbsReMrmj8sehw0Vjocu3sUx9ma4/DV4CuNsmN0mCrD0J5mo25lNgqRCwW5ELmxEk0yqTmTlXY+f/fT2fkpxRwjhBmWFKvYB/K2YXcbA2Q+mMwpWWkpdRvvDkacYRyCCnv9BIowHkeerKyug76OpZCPuhSKgCqMFsqjqgqMXwErABmIw0MvhoMS+k938mTEnxD1CTm53anu1QurTRzx18gHLuK07r6GPNvvN1L5MO4scM7VSo+qzng4HVAzgc11TIL7xQlVNpJ5q1XKdb1Xxs2UzBpjtEXOIh+V9ybPstMl40+gCnTIDtSGXkxPFPNiDaQW3GqkVXBwxEjmV9rW2DkpOCgXah7z/X7zkazP08mrbC7PsrZt01I1qbZlNvi5jJVGnpynk7TytQy3ANjaXa9mMdsOrGtZWYJNhc6CSYaUCI/9pqeXsRCaUDx5Ef4kPU8nJ5anZz+Hi0g7XzO1h3T7TvkELRku0Td6ux3G/x81/+GjZpgsDy8+M5IJFW4jPDvdIDIPdIyDtzlStUgoygludd2SObi3su9x+bkBi7K2SOiaWYFAg6ol44SjLj3BBkmN9J6gtODZYbKJI/5Gd/tk9LjgHIz/W9t9pby5nt3RhC6HN1mtC/SxrMUyWEtzGt506B2fCLjWUclU2bASbWNMVBs2qOJOhxBSMv7BqsYttdlD+FbHYiH4i2Uddem6qHJ9v7WPW1/12IpntEYKF33f/wX7ETKa",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},b=void 0,j={},y=[];function k(e){const t={code:"code",p:"p",...(0,i.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Create a New Role"}),"\n",(0,n.jsx)(a(),{method:"post",path:"/user-svc/role",context:"endpoint"}),"\n",(0,n.jsxs)(t.p,{children:["Create a new role.\n",(0,n.jsx)("b",{children:"The role ID must be prefixed by the caller's slug."}),"\nEg. if the caller's slug is ",(0,n.jsx)(t.code,{children:"petstore-svc"})," the role should look like ",(0,n.jsx)(t.code,{children:"petstore-svc:admin"}),".\nThe user account who creates the role will become the owner of that role, and only the owner will be able to edit the role."]}),"\n",(0,n.jsxs)(t.p,{children:["Requires the ",(0,n.jsx)(t.code,{children:"user-svc:role:create"})," permission."]}),"\n",(0,n.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(c(),{parameters:void 0}),"\n",(0,n.jsx)(p(),{title:"Body",body:{content:{"application/json":{schema:{properties:{description:{type:"string"},id:{type:"string"},name:{type:"string"},permissionIds:{items:{type:"string"},type:"array"}},required:["id"],type:"object"}}},description:"Create Role Request",required:!0}}),"\n",(0,n.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"Role created successfully",content:{"application/json":{schema:{properties:{role:{properties:{createdAt:{type:"string"},description:{type:"string"},id:{type:"string"},name:{type:"string"},ownerId:{type:"string"},updatedAt:{type:"string"}},type:"object"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function v(e={}){const{wrapper:t}={...(0,i.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(k,{...e})}):k(e)}}}]);