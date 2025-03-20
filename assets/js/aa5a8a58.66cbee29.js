"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1269],{40127:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>j,contentTitle:()=>g,default:()=>N,frontMatter:()=>h,metadata:()=>i,toc:()=>y});const i=JSON.parse('{"id":"1backend/list-grants","title":"List Grants","description":"List grants.","source":"@site/docs/1backend/list-grants.api.mdx","sourceDirName":"1backend","slug":"/1backend/list-grants","permalink":"/docs/1backend/list-grants","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"list-grants","title":"List Grants","description":"List grants.","sidebar_label":"List Grants","hide_title":true,"hide_table_of_contents":true,"api":"eJzVVk1v2zgQ/SvEnBXJSbbArk6bLIoiaIAEcXNKDHRMjSU2FMmQlF2voP++GMpK3NjbFj0ssBebpubjzbyZJ/dQUZBeuaisgRKuVYii9mhiyB/No/mQjqKilTIkNo2SjQi6q4NATwJDULWhSgRHUq2UFI58q0JQ1oRM2DV5ryplahEb4iDY6SikNStVdx45Z8pyR8+d8hSS2ecukD8Ja1kmHOVa0ebzXuAcMrCORverCkrQKsQRKGTg6bmjEC9ttYWyB2lNJBP5iM5pJZNb8SVwuT0E2VCLfHKeg0ZFIf16SccZeohbR1BCiF6ZGoYMuAlHHgzZdGOXX0hGGPjqSI93jb0b0e5wK08VlNF3NPBFcNaEEc/ZbMZf3wa6+QjZrxY4cswnFakNhwbqeN0/1ZhwiHWehmbT2DQ4KTtVie/b/YjZK5yDyLsL9B63xzr9Y4tExm+z00N49wa72Fiv/qbq15tK3lv/k1MxZPDuGKlXJpI3qMWc/Jq8eJ9i/jeQmD6SnVdxC+VDD5eEnvxFFxsoHxbDIoOITO4D3AfyYr6WsMigpdhY3kNn0yQ7ZHsopj0u6mk3QyoppNid12xUaCtRNzbE8t3vZ+enwFkmEHMuaqxjH8rbln3aOhKPO5NHECurtd1QJZZbgSI4lCTQVCLaJzIC5bhoYuVtmyZwKkZc21oZQaZyVpnIQqM4fkNYEXNgsOWOXexGJTEAr2Pn1Eca545ZuXsVovdfsXWaDoVlomPSk71pV2ZlJwFDmXinFhU3LaCm8GdQpu40Rm9NLm27B+/2Ssw756xnNsY+NzG6sihOlyifyFTsUMCBMl1cnRiMak2iVdJbpktJCsJpjCvrW+6IVpJMSLVM+T7cXov1eT77Jlsoi2Kz2eS16XLr62LnFwqsnT45z2d5E1udtpZ8G25W8zHbK9iwwbomnytbJJOCW60i9xFOL8dCIAOeqBH+LD/PZyde5md/JKmyIbZo9pDuKS+8qX3vTfF/eQnuxi7S11g4jcpw0YmAfreDDzCFgWxS/EUGvG38sO+XGOje62Hg6+eOPO/9IoM1eoVL7vTDYsimBeC1faItlPDX2KsT3jymAHU3bsAbYRqyyeNCSnLxu7b7UnJ7M/8EGSx3b/HWVuzjccNvStxACelfAHsneUh3PWg0dYc1244xeRlxJxqva8qQsunAVU2PzHYP4ds1HwvhTy7rqEvfjyIwDC/246N/9XjRltGaSVwMw/APurRXUA==","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Change User Password","permalink":"/docs/1backend/change-password"},"next":{"title":"Save Grants","permalink":"/docs/1backend/save-grants"}}');var n=t(74848),a=t(28453),r=t(53746),o=t.n(r),d=t(56518),p=t.n(d),c=t(99972),l=t.n(c),b=t(25342),u=t.n(b),m=(t(44215),t(82223),t(24861));const h={id:"list-grants",title:"List Grants",description:"List grants.",sidebar_label:"List Grants",hide_title:!0,hide_table_of_contents:!0,api:"eJzVVk1v2zgQ/SvEnBXJSbbArk6bLIoiaIAEcXNKDHRMjSU2FMmQlF2voP++GMpK3NjbFj0ssBebpubjzbyZJ/dQUZBeuaisgRKuVYii9mhiyB/No/mQjqKilTIkNo2SjQi6q4NATwJDULWhSgRHUq2UFI58q0JQ1oRM2DV5ryplahEb4iDY6SikNStVdx45Z8pyR8+d8hSS2ecukD8Ja1kmHOVa0ebzXuAcMrCORverCkrQKsQRKGTg6bmjEC9ttYWyB2lNJBP5iM5pJZNb8SVwuT0E2VCLfHKeg0ZFIf16SccZeohbR1BCiF6ZGoYMuAlHHgzZdGOXX0hGGPjqSI93jb0b0e5wK08VlNF3NPBFcNaEEc/ZbMZf3wa6+QjZrxY4cswnFakNhwbqeN0/1ZhwiHWehmbT2DQ4KTtVie/b/YjZK5yDyLsL9B63xzr9Y4tExm+z00N49wa72Fiv/qbq15tK3lv/k1MxZPDuGKlXJpI3qMWc/Jq8eJ9i/jeQmD6SnVdxC+VDD5eEnvxFFxsoHxbDIoOITO4D3AfyYr6WsMigpdhY3kNn0yQ7ZHsopj0u6mk3QyoppNid12xUaCtRNzbE8t3vZ+enwFkmEHMuaqxjH8rbln3aOhKPO5NHECurtd1QJZZbgSI4lCTQVCLaJzIC5bhoYuVtmyZwKkZc21oZQaZyVpnIQqM4fkNYEXNgsOWOXexGJTEAr2Pn1Eca545ZuXsVovdfsXWaDoVlomPSk71pV2ZlJwFDmXinFhU3LaCm8GdQpu40Rm9NLm27B+/2Ssw756xnNsY+NzG6sihOlyifyFTsUMCBMl1cnRiMak2iVdJbpktJCsJpjCvrW+6IVpJMSLVM+T7cXov1eT77Jlsoi2Kz2eS16XLr62LnFwqsnT45z2d5E1udtpZ8G25W8zHbK9iwwbomnytbJJOCW60i9xFOL8dCIAOeqBH+LD/PZyde5md/JKmyIbZo9pDuKS+8qX3vTfF/eQnuxi7S11g4jcpw0YmAfreDDzCFgWxS/EUGvG38sO+XGOje62Hg6+eOPO/9IoM1eoVL7vTDYsimBeC1faItlPDX2KsT3jymAHU3bsAbYRqyyeNCSnLxu7b7UnJ7M/8EGSx3b/HWVuzjccNvStxACelfAHsneUh3PWg0dYc1244xeRlxJxqva8qQsunAVU2PzHYP4ds1HwvhTy7rqEvfjyIwDC/246N/9XjRltGaSVwMw/APurRXUA==",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},g=void 0,j={},y=[];function x(e){const s={code:"code",p:"p",...(0,a.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"List Grants"}),"\n",(0,n.jsx)(o(),{method:"post",path:"/user-svc/grants",context:"endpoint"}),"\n",(0,n.jsx)(s.p,{children:"List grants."}),"\n",(0,n.jsx)(s.p,{children:"Grants define which slugs are assigned specific permissions, overriding the default configuration."}),"\n",(0,n.jsxs)(s.p,{children:["Requires the ",(0,n.jsx)(s.code,{children:"user-svc:grant:view"})," permission."]}),"\n",(0,n.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(p(),{parameters:void 0}),"\n",(0,n.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{permissionId:{type:"string"},slug:{type:"string"}},type:"object"}}},description:"List Grants Request",required:!0}}),"\n",(0,n.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{grants:{items:{properties:{id:{type:"string"},permissionId:{type:"string"},slugs:{description:"Slugs who are granted the PermissionId",items:{type:"string"},type:"array"}},type:"object"},type:"array"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function N(e={}){const{wrapper:s}={...(0,a.R)(),...e.components};return s?(0,n.jsx)(s,{...e,children:(0,n.jsx)(x,{...e})}):x(e)}}}]);