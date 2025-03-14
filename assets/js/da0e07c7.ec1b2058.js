"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5894],{39120:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>R,contentTitle:()=>v,default:()=>f,frontMatter:()=>h,metadata:()=>r,toc:()=>j});const r=JSON.parse('{"id":"1backend/remove-secrets","title":"Remove Secrets","description":"Remove secrets if authorized to do so","source":"@site/docs/1backend/remove-secrets.api.mdx","sourceDirName":"1backend","slug":"/1backend/remove-secrets","permalink":"/docs/1backend/remove-secrets","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"remove-secrets","title":"Remove Secrets","description":"Remove secrets if authorized to do so","sidebar_label":"Remove Secrets","hide_title":true,"hide_table_of_contents":true,"api":"eJylVcFu4zYQ/RVizrLkJF2g1alOaxRBA3QRb05eH2hqLHFDkVySslcVBPQj+oX9kmJEKXISd7FtLxJFznBm3rw36qBAL5y0QRoNOTxgbY7IPAqHwTN5YLwJlXHydyxYMKwwzBtIwFh0nHzuCsjBDV6b6AQJOPzcoA+3pmgh70AYHVAHWnJrlRSDZ/bJU8gOvKiw5rSyju4NEj19yYKeobUIOfjgpC6hT0AW8TRg7S8ajBvcOd7S9xO2F+2esP1XN80bZv8JRYCeti7iF6FgDxGHERHpsIA8uAZ72vDWaB8rvV4u6fX1q6I9JN8O55ts+wS+W169DfWo5y7/hwATYEOAd5dqudMBneaKbdAd0bG1c8b9v0h9Ah5F42RoId92cIvcoVs1oYJ8u+t3CQReesi3MCK4OQrYJVBjqAyRtkCFgfC0nHwgi6Rf+KMYl0RlPyTshxCNU2SXKSO4qowP+bvvr2+ugIJNuWwo5djV84xeA/Khtcg+jiYfgR2MUuaEBdu3jDNvuUDGNUnuCTXjIrKHHZypWaiQPXp0VBG7N6XUDHVhjdQhhQQk3V8hL5AQ1rwm4FZjfwd8YWa2lb9ipDZh/jDrdv2F11bhpMMJ+VF+2+l7N8prPo+qmg1IsfpgpjnAxdBtrLkkMD1X6H/0UpeN4sEZnQpTn6X9/o5tGmuNIxFF/KsQbJ5lV3sunlAX5JDBGyGumDB1bTQbzdjBONaaxrHVHTvjmv/rjz9rKZyhRkuBfrHnHosB/H0jVaCp5wVXSNgqKZBEmHfPGf7y/p4db9Lli/x8nmWn0yktdZMaV2ajn894adXiJl2mVajVMGLQ1f63wyZGn8vzJ16W6FJpssEko6bJQB2Bq9tYEyRA3IwFL9ObdLlwIr3+ge61xoea67NMX8wTovYLwM6m9Df/BEYSBfwSMqu41BR4AKEbRbWFWVSDmGLsXQIkHzrvOsL70am+p+3PDTrS8y6BI3eS76ngLZFoZDTpMDLup5jvgqRESHDVREq/miN9MnmshEAbvmp7PiB+Xt+vP6whgf34H6tNQV6On2ii8xPkMPwKBx6RwbDXgeK6bHhJtvFW0hcf58CsPEoqmRZU13Sk27McXys3lkJPKuyiS9dFXff9s308+keP53ERramTu77v/wZuF8kk","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Check Security Status","permalink":"/docs/1backend/is-secure"},"next":{"title":"List Secrets","permalink":"/docs/1backend/list-secrets"}}');var n=s(74848),o=s(28453),i=s(53746),a=s.n(i),d=s(56518),c=s.n(d),p=s(99972),l=s.n(p),u=s(25342),m=s.n(u),b=(s(44215),s(82223),s(24861));const h={id:"remove-secrets",title:"Remove Secrets",description:"Remove secrets if authorized to do so",sidebar_label:"Remove Secrets",hide_title:!0,hide_table_of_contents:!0,api:"eJylVcFu4zYQ/RVizrLkJF2g1alOaxRBA3QRb05eH2hqLHFDkVySslcVBPQj+oX9kmJEKXISd7FtLxJFznBm3rw36qBAL5y0QRoNOTxgbY7IPAqHwTN5YLwJlXHydyxYMKwwzBtIwFh0nHzuCsjBDV6b6AQJOPzcoA+3pmgh70AYHVAHWnJrlRSDZ/bJU8gOvKiw5rSyju4NEj19yYKeobUIOfjgpC6hT0AW8TRg7S8ajBvcOd7S9xO2F+2esP1XN80bZv8JRYCeti7iF6FgDxGHERHpsIA8uAZ72vDWaB8rvV4u6fX1q6I9JN8O55ts+wS+W169DfWo5y7/hwATYEOAd5dqudMBneaKbdAd0bG1c8b9v0h9Ah5F42RoId92cIvcoVs1oYJ8u+t3CQReesi3MCK4OQrYJVBjqAyRtkCFgfC0nHwgi6Rf+KMYl0RlPyTshxCNU2SXKSO4qowP+bvvr2+ugIJNuWwo5djV84xeA/Khtcg+jiYfgR2MUuaEBdu3jDNvuUDGNUnuCTXjIrKHHZypWaiQPXp0VBG7N6XUDHVhjdQhhQQk3V8hL5AQ1rwm4FZjfwd8YWa2lb9ipDZh/jDrdv2F11bhpMMJ+VF+2+l7N8prPo+qmg1IsfpgpjnAxdBtrLkkMD1X6H/0UpeN4sEZnQpTn6X9/o5tGmuNIxFF/KsQbJ5lV3sunlAX5JDBGyGumDB1bTQbzdjBONaaxrHVHTvjmv/rjz9rKZyhRkuBfrHnHosB/H0jVaCp5wVXSNgqKZBEmHfPGf7y/p4db9Lli/x8nmWn0yktdZMaV2ajn894adXiJl2mVajVMGLQ1f63wyZGn8vzJ16W6FJpssEko6bJQB2Bq9tYEyRA3IwFL9ObdLlwIr3+ge61xoea67NMX8wTovYLwM6m9Df/BEYSBfwSMqu41BR4AKEbRbWFWVSDmGLsXQIkHzrvOsL70am+p+3PDTrS8y6BI3eS76ngLZFoZDTpMDLup5jvgqRESHDVREq/miN9MnmshEAbvmp7PiB+Xt+vP6whgf34H6tNQV6On2ii8xPkMPwKBx6RwbDXgeK6bHhJtvFW0hcf58CsPEoqmRZU13Sk27McXys3lkJPKuyiS9dFXff9s308+keP53ERramTu77v/wZuF8kk",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},v=void 0,R={},j=[];function y(e){const t={p:"p",...(0,o.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Remove Secrets"}),"\n",(0,n.jsx)(a(),{method:"delete",path:"/secret-svc/secrets",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Remove secrets if authorized to do so"}),"\n",(0,n.jsx)(b.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(c(),{parameters:void 0}),"\n",(0,n.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{id:{type:"string"},ids:{items:{type:"string"},type:"array"},key:{type:"string"},keys:{items:{type:"string"},type:"array"}},type:"object"}}},description:"Remove Secret Request",required:!0}}),"\n",(0,n.jsx)(m(),{id:void 0,label:void 0,responses:{200:{description:"Remove Secret Response",content:{"application/json":{schema:{type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function f(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(y,{...e})}):y(e)}}}]);