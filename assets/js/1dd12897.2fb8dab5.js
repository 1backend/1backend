"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[8211],{68361:(e,i,s)=>{s.r(i),s.d(i,{assets:()=>F,contentTitle:()=>f,default:()=>m,frontMatter:()=>h,metadata:()=>t,toc:()=>g});const t=JSON.parse('{"id":"1backend/serve-upload","title":"Serve an Uploaded File","description":"Retrieves and serves a previously uploaded file using its File ID.","source":"@site/docs/1backend/serve-upload.api.mdx","sourceDirName":"1backend","slug":"/1backend/serve-upload","permalink":"/docs/1backend/serve-upload","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"serve-upload","title":"Serve an Uploaded File","description":"Retrieves and serves a previously uploaded file using its File ID.","sidebar_label":"Serve an Uploaded File","hide_title":true,"hide_table_of_contents":true,"api":"eJztVlFv20YM/ivEPbWALDlJC2x6Wop0hbFuK+rkKQ6Q84myuJzulOPJnmoI2I/YL9wvGShZdZYlGzBgBQbsyZaO5H38yI/UXhXIJlATyTuVq48YA+EWGbQrgDEMf6EJuCXfsu2gbazXBRZQkkVomdwGKDJ8K4+Li3TlfvARc7isEG4XF7dDoFs5lYeS0BYMvgTtDqFAB4SCyhIDupiu3OxoTnJ56+i+RaACXaSSMEDpA8QKRwgUGW05+v2NjwZu0FBJBgI2lowWJFOkdOWW5AzCyRtt7tAVY6iCOAZatxEL4I4j1slgzmC0gzVOoeRcm+CZoW5tpMYiOF8gpyt3WRFDjdoxoDbVIxJr3UGlt3j0C2h8KBh2FKsBH+saj7Ss23hkbMiaU5Uo32DQUshFoXI1FO9quEglqtFB1xgxsMqv94+qPsY9kGa7I238mGZ44XyU2i0uEthVZCpJPyCji/wEvy9VokiuaHSsVKKcrlHlSiIuBFbA+5YCFiqPocVEsamw1irfq9g1Yincu43q+xsx5sY7Rpbz0/lcfv6cyNi1BXBrDDKXrbWdSpTxLqKL4qObsWDkXeZNxDjjGFDXcvYcgESVPtQ6qlytyenQqb7v+0S9egrG98SDLg4Nvrj4J/c3QeoZaUwXQ/DhKV6S6Y1f/4Qmfob16hl2pH6lb13xxTG9foqqhYsYnLawlKoFeDvE/LLQ5KXeiDBGipZbo24SVWOsvChpg3FQUKxUrjLp3RlvTTb0WTbWONuPLd2rZBTeQWdtsOKTWW+0rTzH/PVXp2cnStqZ0bSBYreUHEbYb1AHDOet3PSYqcuuQVgdTFYKSm+t32EB624QnjY4zNro79CBNqOwoAy+HlR8xRgkNXjvN+QAXdF4cjGdFFqhLjAcNSoofKBPA/HqM2+6oe+wG5gkV3rBKbXSZqgV1pokY9YW+RsRQWt1DN6lxtcPYn9YwLJtGh+E2ZGkKsYmz7KT9Th8xSFTffKIhXOZUiVJ2ochLYO9822A8wU8aBX+7Zdfa5J5jGFLBnm21iwj2hWwbslGiB7YaIvCgCWDjlFSmCC++/Aetmfp/A8AOc+y3W6Xblyb+rDJDn6c6U1jZ2fpPK1ibQV2xFDzj+VyvP2YH+/0ZoMhJZ8NJplQS9GKybR4VKKkg8aM5+lZOp8Fk55+LXEbz7HW7gHSQTkyk6+mrSJdrB4xtz9K6v8t/1/e8gcdRvw5Zo3V5KQrhg7dH2bUtZpm1DSNpIenL4H8sHxvEiUDScz3e9HGVbB9L6/vWwydyq9vErXVgfRamlM+Gojlf6HyUlvGv2iwFx8Pe/0l/LvfFk+SMU0qJ3t/q20rTypRd9gdvz36mz6ZZp4kNx6eG4NNfOD27PLpH26Id28vVd//Dj2B8E8=","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Serve a Downloaded file","permalink":"/docs/1backend/serve-download"},"next":{"title":"Upload a File","permalink":"/docs/1backend/upload-file"}}');var l=s(74848),a=s(28453),o=s(53746),d=s.n(o),r=s(56518),n=s.n(r),c=s(99972),p=s.n(c),u=s(25342),v=s.n(u),b=(s(44215),s(82223),s(24861));const h={id:"serve-upload",title:"Serve an Uploaded File",description:"Retrieves and serves a previously uploaded file using its File ID.",sidebar_label:"Serve an Uploaded File",hide_title:!0,hide_table_of_contents:!0,api:"eJztVlFv20YM/ivEPbWALDlJC2x6Wop0hbFuK+rkKQ6Q84myuJzulOPJnmoI2I/YL9wvGShZdZYlGzBgBQbsyZaO5H38yI/UXhXIJlATyTuVq48YA+EWGbQrgDEMf6EJuCXfsu2gbazXBRZQkkVomdwGKDJ8K4+Li3TlfvARc7isEG4XF7dDoFs5lYeS0BYMvgTtDqFAB4SCyhIDupiu3OxoTnJ56+i+RaACXaSSMEDpA8QKRwgUGW05+v2NjwZu0FBJBgI2lowWJFOkdOWW5AzCyRtt7tAVY6iCOAZatxEL4I4j1slgzmC0gzVOoeRcm+CZoW5tpMYiOF8gpyt3WRFDjdoxoDbVIxJr3UGlt3j0C2h8KBh2FKsBH+saj7Ss23hkbMiaU5Uo32DQUshFoXI1FO9quEglqtFB1xgxsMqv94+qPsY9kGa7I238mGZ44XyU2i0uEthVZCpJPyCji/wEvy9VokiuaHSsVKKcrlHlSiIuBFbA+5YCFiqPocVEsamw1irfq9g1Yincu43q+xsx5sY7Rpbz0/lcfv6cyNi1BXBrDDKXrbWdSpTxLqKL4qObsWDkXeZNxDjjGFDXcvYcgESVPtQ6qlytyenQqb7v+0S9egrG98SDLg4Nvrj4J/c3QeoZaUwXQ/DhKV6S6Y1f/4Qmfob16hl2pH6lb13xxTG9foqqhYsYnLawlKoFeDvE/LLQ5KXeiDBGipZbo24SVWOsvChpg3FQUKxUrjLp3RlvTTb0WTbWONuPLd2rZBTeQWdtsOKTWW+0rTzH/PVXp2cnStqZ0bSBYreUHEbYb1AHDOet3PSYqcuuQVgdTFYKSm+t32EB624QnjY4zNro79CBNqOwoAy+HlR8xRgkNXjvN+QAXdF4cjGdFFqhLjAcNSoofKBPA/HqM2+6oe+wG5gkV3rBKbXSZqgV1pokY9YW+RsRQWt1DN6lxtcPYn9YwLJtGh+E2ZGkKsYmz7KT9Th8xSFTffKIhXOZUiVJ2ochLYO9822A8wU8aBX+7Zdfa5J5jGFLBnm21iwj2hWwbslGiB7YaIvCgCWDjlFSmCC++/Aetmfp/A8AOc+y3W6Xblyb+rDJDn6c6U1jZ2fpPK1ibQV2xFDzj+VyvP2YH+/0ZoMhJZ8NJplQS9GKybR4VKKkg8aM5+lZOp8Fk55+LXEbz7HW7gHSQTkyk6+mrSJdrB4xtz9K6v8t/1/e8gcdRvw5Zo3V5KQrhg7dH2bUtZpm1DSNpIenL4H8sHxvEiUDScz3e9HGVbB9L6/vWwydyq9vErXVgfRamlM+Gojlf6HyUlvGv2iwFx8Pe/0l/LvfFk+SMU0qJ3t/q20rTypRd9gdvz36mz6ZZp4kNx6eG4NNfOD27PLpH26Id28vVd//Dj2B8E8=",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,F={},g=[];function D(e){const i={code:"code",li:"li",p:"p",ul:"ul",...(0,a.R)(),...e.components};return(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Serve an Uploaded File"}),"\n",(0,l.jsx)(d(),{method:"get",path:"/file-svc/serve/upload/{fileId}",context:"endpoint"}),"\n",(0,l.jsxs)(i.p,{children:["Retrieves and serves a previously uploaded file using its File ID.\nNote: The ",(0,l.jsx)(i.code,{children:"ID"})," and ",(0,l.jsx)(i.code,{children:"FileID"})," fields of an upload are different."]}),"\n",(0,l.jsxs)(i.ul,{children:["\n",(0,l.jsxs)(i.li,{children:[(0,l.jsx)(i.code,{children:"FileID"})," is a unique identifier for the file itself."]}),"\n",(0,l.jsxs)(i.li,{children:[(0,l.jsx)(i.code,{children:"ID"})," is a unique identifier for a specific replica of the file.\nSince 1Backend is a distributed system, files can be replicated across multiple nodes.\nThis means each uploaded file may have multiple records with the same ",(0,l.jsx)(i.code,{children:"FileID"})," but different ",(0,l.jsx)(i.code,{children:"ID"}),"s."]}),"\n"]}),"\n",(0,l.jsx)(b.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,l.jsx)(n(),{parameters:[{description:"FileID uniquely identifies the file itself (not an ID, which represents a specific replica)",in:"path",name:"fileId",required:!0,schema:{type:"string"}}]}),"\n",(0,l.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,l.jsx)(v(),{id:void 0,label:void 0,responses:{200:{description:"File served successfully",content:{"application/octet-stream":{schema:{type:"string",format:"binary"}}}},400:{description:"Missing upload ID",content:{"application/octet-stream":{schema:{properties:{error:{type:"string"}},type:"object"}}}},404:{description:"File not found",content:{"application/octet-stream":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/octet-stream":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function m(e={}){const{wrapper:i}={...(0,a.R)(),...e.components};return i?(0,l.jsx)(i,{...e,children:(0,l.jsx)(D,{...e})}):D(e)}}}]);