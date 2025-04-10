"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6571],{32109:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>j,contentTitle:()=>m,default:()=>y,frontMatter:()=>f,metadata:()=>s,toc:()=>h});const s=JSON.parse('{"id":"1backend/is-secure","title":"Check Security Status","description":"Returns true if the encryption key is sufficiently secure.","source":"@site/docs/1backend/is-secure.api.mdx","sourceDirName":"1backend","slug":"/1backend/is-secure","permalink":"/docs/1backend/is-secure","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"is-secure","title":"Check Security Status","description":"Returns true if the encryption key is sufficiently secure.","sidebar_label":"Check Security Status","hide_title":true,"hide_table_of_contents":true,"api":"eJylVMuO2zgQ/BWiz7JkrxMsoFNmFoNgsAESjDK5OD60qbbEmCI5JGWvIujfg5b8ipMcFnOxBbG6q7pYrR5KCtIrF5U1kMMTxdabIKJvSaitiDUJMtJ3I0DsqBMqiNBut0oqMlF3IpBsPaWQgHXkkXGPJeSgQjGeQAKegrMmUIC8h7/mc/77mfdh4hBfULckno54SEBaE8lErkDntJIjQfYtcFkPQdbUID85z/RRTSRn8ryH2DmCHDbWakIDw8CCXlrlqYR8dYGukxPUbr6RjDAMjH3zO733WIonemkpxP8j8tg/RK9Mde6/+LX/s8E21tar71S+luDt7wZ4NJG8QS0K8nvy4sF761/HNCQwZkHFDvJVD/eEnvxdG2vIV+uB7cUqsOUFSU9RFHvJpjcUa8uJqYjNdMgFkIURNAt7makwC6cshVFvGBlarxmZaStR1zbEfLFYLv8G5jpJKVjxFIprQbd+fO4cia9HyFcQW6u1PVApNp1AERxKEmhKEe2OjEA55UdsvW3GJXkO5Hkg8cFWyggypbPKRF4Lxf1rwpLYYIMN+3Z3vN7RXjgnD536l7oxo8psLevkG0E53gg1qHjigJrCu6BM1WqM3ppU2uaq96dHUbTOWc9+TibVMbo8yxYblDsyJRdkMCQ3Ltw9zgxGtSfRKOkte60kBeE0xq31DY+jlSTezbw/873/9EHsl+n8J7aQZ9nhcEgr06bWV9mxLmRYOT1bpvO0jo1mDZF8Ez5ui4ntIjYcsKrIp8pmIyRjn1TUDFncT4NAAhyHSf48XabzmZfp8i33dTbEBs2V0n9qkjtRHKMhioixDXDjQn9Zgtd9D493Gum/mDmNyrCo0aD+mPIVXFLOSTnnfJ0A55kRfb/BQM9eDwO/fmnJ836tE9ijV7hhO1brITlFjBdjRx3fi5TkOAJ7/qqO6brZ6OF6/d4/fIYE8Lgel0Bys+T0wN1PR6a76n0b6EkC/w7JH0r6for7MJzx09EfK85bNKHZ0fUwDD8APkxTfw==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Encrypt a Value","permalink":"/docs/1backend/encrypt-value"},"next":{"title":"Remove Secrets","permalink":"/docs/1backend/remove-secrets"}}');var i=n(74848),r=n(28453),c=n(53746),a=n.n(c),o=n(56518),d=n.n(o),p=n(99972),u=n.n(p),l=n(25342),b=n.n(l),k=(n(44215),n(82223),n(24861));const f={id:"is-secure",title:"Check Security Status",description:"Returns true if the encryption key is sufficiently secure.",sidebar_label:"Check Security Status",hide_title:!0,hide_table_of_contents:!0,api:"eJylVMuO2zgQ/BWiz7JkrxMsoFNmFoNgsAESjDK5OD60qbbEmCI5JGWvIujfg5b8ipMcFnOxBbG6q7pYrR5KCtIrF5U1kMMTxdabIKJvSaitiDUJMtJ3I0DsqBMqiNBut0oqMlF3IpBsPaWQgHXkkXGPJeSgQjGeQAKegrMmUIC8h7/mc/77mfdh4hBfULckno54SEBaE8lErkDntJIjQfYtcFkPQdbUID85z/RRTSRn8ryH2DmCHDbWakIDw8CCXlrlqYR8dYGukxPUbr6RjDAMjH3zO733WIonemkpxP8j8tg/RK9Mde6/+LX/s8E21tar71S+luDt7wZ4NJG8QS0K8nvy4sF761/HNCQwZkHFDvJVD/eEnvxdG2vIV+uB7cUqsOUFSU9RFHvJpjcUa8uJqYjNdMgFkIURNAt7makwC6cshVFvGBlarxmZaStR1zbEfLFYLv8G5jpJKVjxFIprQbd+fO4cia9HyFcQW6u1PVApNp1AERxKEmhKEe2OjEA55UdsvW3GJXkO5Hkg8cFWyggypbPKRF4Lxf1rwpLYYIMN+3Z3vN7RXjgnD536l7oxo8psLevkG0E53gg1qHjigJrCu6BM1WqM3ppU2uaq96dHUbTOWc9+TibVMbo8yxYblDsyJRdkMCQ3Ltw9zgxGtSfRKOkte60kBeE0xq31DY+jlSTezbw/873/9EHsl+n8J7aQZ9nhcEgr06bWV9mxLmRYOT1bpvO0jo1mDZF8Ez5ui4ntIjYcsKrIp8pmIyRjn1TUDFncT4NAAhyHSf48XabzmZfp8i33dTbEBs2V0n9qkjtRHKMhioixDXDjQn9Zgtd9D493Gum/mDmNyrCo0aD+mPIVXFLOSTnnfJ0A55kRfb/BQM9eDwO/fmnJ836tE9ijV7hhO1brITlFjBdjRx3fi5TkOAJ7/qqO6brZ6OF6/d4/fIYE8Lgel0Bys+T0wN1PR6a76n0b6EkC/w7JH0r6for7MJzx09EfK85bNKHZ0fUwDD8APkxTfw==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,j={},h=[];function D(e){const t={p:"p",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(k.default,{as:"h1",className:"openapi__heading",children:"Check Security Status"}),"\n",(0,i.jsx)(a(),{method:"get",path:"/secret-svc/is-secure",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Returns true if the encryption key is sufficiently secure."}),"\n",(0,i.jsx)(d(),{parameters:void 0}),"\n",(0,i.jsx)(u(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"Encrypt Value Response",content:{"application/json":{schema:{properties:{isSecure:{type:"boolean"}},required:["isSecure"],type:"object"}}}},400:{description:"Bad Request",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function y(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(D,{...e})}):D(e)}}}]);