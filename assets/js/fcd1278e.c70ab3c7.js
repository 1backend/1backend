"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3663],{43861:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>j,contentTitle:()=>f,default:()=>u,frontMatter:()=>b,metadata:()=>a,toc:()=>D});const a=JSON.parse('{"id":"1backend/delete-message","title":"Delete a Message","description":"Delete a specific message from a chat thread by its ID","source":"@site/docs/1backend/delete-message.api.mdx","sourceDirName":"1backend","slug":"/1backend/delete-message","permalink":"/docs/1backend/delete-message","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"delete-message","title":"Delete a Message","description":"Delete a specific message from a chat thread by its ID","sidebar_label":"Delete a Message","hide_title":true,"hide_table_of_contents":true,"api":"eJylVE2P2zYQ/SvEnFqAK9l1iwI6dZNdFG63zaLOnjY+jKmxxCxFKiRl1xX434ux5I91nAJJL7Ykzsx78/hmeigpKK/bqJ2FAu7IUCSBIrSk9For0VAIWJFYe9cIFKrGKGLtCUux2gkdg5jfgQTXkkcuMi+hgHJf5o8hFSS06LGhSD5A8dxfYI5hQx3NX1qMNUiw2BAUMDKYlyDB06dOeyqhiL4jCUHV1CAUPcRdy8Ehem0rSGnJwaF1NlDg8x8mE/67Dh06pSiEdWfMTgzkGU05G8lGzsO2NVrtO8w/Bk7uz8CxLDUfoXn0rETUDDpQHIm51UdSEVJKScKP18jM7QaNLsVvi3d/fg34ZecDwPRzgCeLXayd1/98XXfXAH663kEkb9GIBfkNeXHvvfP/DylJCKQ6r+Nub503hJ78bRdrKJ6XfMsRK3YVvGVnLjYKlhIairU7GXHvQM6AnP17EzYqH12V90d7JWAsJj64tPOGM3LjFJrahVhMp7PZz8CgB04Lpj4Y7JzZpTDvdy2JD2PIBxBrZ4zb0n6EeNZQkUBbiuheyApUg8mHmYs1iadAnnsTD67SVpAtW6dtzA4DUxOW5E8jczve815nOHoQW/077YBF1XbtmCdfDar91VCDmjsOaCj8ErStOoPRO5sp15zVfpyLRde2zkeQo0h1jG2R59MVqheyJSfkkOSFCrfzG4tRb0g0WnnHWmtFQbQG49r5htsxWpENxHwOeL8+PojNLJu8QgtFnm+326yyXeZ8lY95IceqNTezbJLVsTHMIZJvwrv1YkA7kQ1brCrymXb5PiRnnXQ0HDJ9MzQCEtgOA/1JNssmN15l7AEJrQuxQXvG9Lg8T4vvlQD9aRC+fdGOVxnp75i3BrVlLntd+tHkz3AwOcjD8gQJxWmPLiWwnzm071cY6MmblPjzp448D9pSwga9xhXLwRtbB34uoVijCfQfjX3317iivxevFvtV3gdf2h0rjabjN5DwQrtXiz8tkzyYnOkM57dKURvPMj9bLul8F9zdP9y/vwcJOM7oaSq4njw8MMBVYpdTNbDg3yS/kNL3w8yldIwfjr6YcRzlIZp1WqaU/gWafpyH","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Events","permalink":"/docs/1backend/events"},"next":{"title":"Get Message","permalink":"/docs/1backend/get-message"}}');var n=s(74848),i=s(28453),d=s(53746),r=s.n(d),l=s(56518),c=s.n(l),o=s(99972),p=s.n(o),g=s(25342),m=s.n(g),h=(s(44215),s(82223),s(24861));const b={id:"delete-message",title:"Delete a Message",description:"Delete a specific message from a chat thread by its ID",sidebar_label:"Delete a Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVE2P2zYQ/SvEnFqAK9l1iwI6dZNdFG63zaLOnjY+jKmxxCxFKiRl1xX434ux5I91nAJJL7Ykzsx78/hmeigpKK/bqJ2FAu7IUCSBIrSk9For0VAIWJFYe9cIFKrGKGLtCUux2gkdg5jfgQTXkkcuMi+hgHJf5o8hFSS06LGhSD5A8dxfYI5hQx3NX1qMNUiw2BAUMDKYlyDB06dOeyqhiL4jCUHV1CAUPcRdy8Ehem0rSGnJwaF1NlDg8x8mE/67Dh06pSiEdWfMTgzkGU05G8lGzsO2NVrtO8w/Bk7uz8CxLDUfoXn0rETUDDpQHIm51UdSEVJKScKP18jM7QaNLsVvi3d/fg34ZecDwPRzgCeLXayd1/98XXfXAH663kEkb9GIBfkNeXHvvfP/DylJCKQ6r+Nub503hJ78bRdrKJ6XfMsRK3YVvGVnLjYKlhIairU7GXHvQM6AnP17EzYqH12V90d7JWAsJj64tPOGM3LjFJrahVhMp7PZz8CgB04Lpj4Y7JzZpTDvdy2JD2PIBxBrZ4zb0n6EeNZQkUBbiuheyApUg8mHmYs1iadAnnsTD67SVpAtW6dtzA4DUxOW5E8jczve815nOHoQW/077YBF1XbtmCdfDar91VCDmjsOaCj8ErStOoPRO5sp15zVfpyLRde2zkeQo0h1jG2R59MVqheyJSfkkOSFCrfzG4tRb0g0WnnHWmtFQbQG49r5htsxWpENxHwOeL8+PojNLJu8QgtFnm+326yyXeZ8lY95IceqNTezbJLVsTHMIZJvwrv1YkA7kQ1brCrymXb5PiRnnXQ0HDJ9MzQCEtgOA/1JNssmN15l7AEJrQuxQXvG9Lg8T4vvlQD9aRC+fdGOVxnp75i3BrVlLntd+tHkz3AwOcjD8gQJxWmPLiWwnzm071cY6MmblPjzp448D9pSwga9xhXLwRtbB34uoVijCfQfjX3317iivxevFvtV3gdf2h0rjabjN5DwQrtXiz8tkzyYnOkM57dKURvPMj9bLul8F9zdP9y/vwcJOM7oaSq4njw8MMBVYpdTNbDg3yS/kNL3w8yldIwfjr6YcRzlIZp1WqaU/gWafpyH",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,j={},D=[];function k(e){const t={p:"p",...(0,i.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Delete a Message"}),"\n",(0,n.jsx)(r(),{method:"delete",path:"/chat-svc/message/{messageId}",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Delete a specific message from a chat thread by its ID"}),"\n",(0,n.jsx)(h.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(c(),{parameters:[{description:"Message ID",in:"path",name:"messageId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(m(),{id:void 0,label:void 0,responses:{200:{description:"Message successfully deleted",content:{"application/json":{schema:{additionalProperties:!0,type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function u(e={}){const{wrapper:t}={...(0,i.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(k,{...e})}):k(e)}}}]);