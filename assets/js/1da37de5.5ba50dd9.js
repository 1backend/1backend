"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6759],{28778:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>S,contentTitle:()=>y,default:()=>f,frontMatter:()=>m,metadata:()=>n,toc:()=>C});const n=JSON.parse('{"id":"1backend/delete-node","title":"Delete Node","description":"Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\'s still present in the database.","source":"@site/docs/1backend/delete-node.api.mdx","sourceDirName":"1backend","slug":"/1backend/delete-node","permalink":"/docs/1backend/delete-node","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"delete-node","title":"Delete Node","description":"Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\'s still present in the database.","sidebar_label":"Delete Node","hide_title":true,"hide_table_of_contents":true,"api":"eJzNVU1v3EYM/SsDXpoCsrS2E6DVqU5tFIsaaeCPk7MHrsSVJh7NTDij3W4E/feCkta7tjc5Gr3o85F85PCRHZQUCtY+amchh0syFCkoVEyVDpGYSmVdSWq5He/3N9epuqt1UGRL77SNSgfVBlq1Rm1qsgpHoA7KOmWcrYgVrlEbXBpSyzYqHX8JKkRtjPJMgcSHVbEmVWLEJQZKIQHniVF4zUvIoRyYfXIlQQIeGRuKxAHyh+5FDp8mmpCAlnePsYYELDYEObRsIAGmb61mKiGP3FICoaipQcg7iFsvsBBZ2wr6fiHg4J0NFOT/2ey93F6GVH86G8lG6BN4P5u9hsztGo0u1fwSEigmcN4Bem90MeSZfQ0C7Q7YeJYqRD3GJmbHx0gmuy9u+ZWKCH3fDzxOX/O4t9jG2rH+TuUbMjlStFvitS5IWRfVyrX27eh8OH5AkdiiUcKLWF0NPt+GUp9AoKJlHbdDQ38kZOKLNtaQPyykByNW0utwM8iSt+p2XcAigYZi7fb6GLQhVpDxhDwJ6yITRWZdy6YHCSUJjtIROeSQZcYVaGoXYv7ht7PzU5CYO0q3kuKY1SGxlwW823pSXybIF1ArZ4zbUCmjA1XwWJBCW6roHmVKFKMC1YpdM2j/PhBLWuraVdo+jZd0p+OasCTeK/li6uPhPOCprOj137QdCq3tyglPOUIshiOkBrVkHNBQ+CNoW7UGIzubFq458P15rm5b7x1HSKYi1TH6PMtOl1g8ki3FIBO5P6/CxfzEYtRrUo0u2IWxyYPyBuPKcSPpGF2QDSR8dvH++nyt1ufp7Fm0kGfZZrNJK9umjqtssgsZVt6cnKeztI6NEQ6RuAn/rCZJ7cmGDVYVcapdNkAyqZOORiCnH8dEIAFph5H+LD1PZydcpGe/i1/vQmzQHjAdN4SaBvGz3Lu9Vv5Xi2RqjEj/xswb1FYyG6rcTWp5gEO1SBuM6eWCWiQgwhBU14nPezZ9L5+/tcQi2EUCa2QttMZ9pIM8l5Cv0AT6SZne3UyL6Fd1sLaOEt61t93KgaFp5Q0SeKTttNb6RZ/sVCI0xj8XRUE+Hti8mmL94Ry5vLq+uruCBHAS+V5W4i/ZPUiAo5ReynJkIdc++YFJ142i7fsn/PjrhxZPs2BES4UWfd//B9OuA8M=","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"List Service Instances","permalink":"/docs/1backend/list-instances"},"next":{"title":"View Self Node","permalink":"/docs/1backend/self-node"}}');var s=i(74848),o=i(28453),d=i(53746),a=i.n(d),r=i(56518),l=i.n(r),p=i(99972),c=i.n(p),u=i(25342),b=i.n(u),h=(i(44215),i(82223),i(24861));const m={id:"delete-node",title:"Delete Node",description:"Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it's still present in the database.",sidebar_label:"Delete Node",hide_title:!0,hide_table_of_contents:!0,api:"eJzNVU1v3EYM/SsDXpoCsrS2E6DVqU5tFIsaaeCPk7MHrsSVJh7NTDij3W4E/feCkta7tjc5Gr3o85F85PCRHZQUCtY+amchh0syFCkoVEyVDpGYSmVdSWq5He/3N9epuqt1UGRL77SNSgfVBlq1Rm1qsgpHoA7KOmWcrYgVrlEbXBpSyzYqHX8JKkRtjPJMgcSHVbEmVWLEJQZKIQHniVF4zUvIoRyYfXIlQQIeGRuKxAHyh+5FDp8mmpCAlnePsYYELDYEObRsIAGmb61mKiGP3FICoaipQcg7iFsvsBBZ2wr6fiHg4J0NFOT/2ey93F6GVH86G8lG6BN4P5u9hsztGo0u1fwSEigmcN4Bem90MeSZfQ0C7Q7YeJYqRD3GJmbHx0gmuy9u+ZWKCH3fDzxOX/O4t9jG2rH+TuUbMjlStFvitS5IWRfVyrX27eh8OH5AkdiiUcKLWF0NPt+GUp9AoKJlHbdDQ38kZOKLNtaQPyykByNW0utwM8iSt+p2XcAigYZi7fb6GLQhVpDxhDwJ6yITRWZdy6YHCSUJjtIROeSQZcYVaGoXYv7ht7PzU5CYO0q3kuKY1SGxlwW823pSXybIF1ArZ4zbUCmjA1XwWJBCW6roHmVKFKMC1YpdM2j/PhBLWuraVdo+jZd0p+OasCTeK/li6uPhPOCprOj137QdCq3tyglPOUIshiOkBrVkHNBQ+CNoW7UGIzubFq458P15rm5b7x1HSKYi1TH6PMtOl1g8ki3FIBO5P6/CxfzEYtRrUo0u2IWxyYPyBuPKcSPpGF2QDSR8dvH++nyt1ufp7Fm0kGfZZrNJK9umjqtssgsZVt6cnKeztI6NEQ6RuAn/rCZJ7cmGDVYVcapdNkAyqZOORiCnH8dEIAFph5H+LD1PZydcpGe/i1/vQmzQHjAdN4SaBvGz3Lu9Vv5Xi2RqjEj/xswb1FYyG6rcTWp5gEO1SBuM6eWCWiQgwhBU14nPezZ9L5+/tcQi2EUCa2QttMZ9pIM8l5Cv0AT6SZne3UyL6Fd1sLaOEt61t93KgaFp5Q0SeKTttNb6RZ/sVCI0xj8XRUE+Hti8mmL94Ry5vLq+uruCBHAS+V5W4i/ZPUiAo5ReynJkIdc++YFJ142i7fsn/PjrhxZPs2BES4UWfd//B9OuA8M=",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},y=void 0,S={},C=[];function H(e){const t={p:"p",...(0,o.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Delete Node"}),"\n",(0,s.jsx)(a(),{method:"delete",path:"/registry-svc/node/{url}",context:"endpoint"}),"\n",(0,s.jsx)(t.p,{children:"Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it's still present in the database."}),"\n",(0,s.jsx)(h.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,s.jsx)(l(),{parameters:[{description:"Node URL",in:"path",name:"url",required:!0,schema:{type:"string"}}]}),"\n",(0,s.jsx)(c(),{title:"Body",body:void 0}),"\n",(0,s.jsx)(b(),{id:void 0,label:void 0,responses:{204:{description:"No Content"},400:{description:"Invalid ID",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},404:{description:"Service not found",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function f(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(H,{...e})}):H(e)}}}]);