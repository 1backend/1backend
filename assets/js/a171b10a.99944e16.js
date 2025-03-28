"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9707],{22972:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>g,contentTitle:()=>h,default:()=>v,frontMatter:()=>j,metadata:()=>i,toc:()=>y});const i=JSON.parse('{"id":"1backend/list-models","title":"List Models","description":"Retrieves a list of models.","source":"@site/docs/1backend/list-models.api.mdx","sourceDirName":"1backend","slug":"/1backend/list-models","permalink":"/docs/1backend/list-models","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"list-models","title":"List Models","description":"Retrieves a list of models.","sidebar_label":"List Models","hide_title":true,"hide_table_of_contents":true,"api":"eJy9Vttu40YM/ZXBPCuSs+kChZ6aFEURbNoEcdIXx8DSEm3PZm7hjOR4Df17QUl2fJFbtCj6Yo95OUNyDklvZImhIOWjclbm8hEjKawxCBBahSjcXBhXog7pi32xj/hWKcIgvrbCi1AXeXvKa4Wrr8IjGRWCcjaViXQeCRj4tpS5ZLjfWiiZSMLgnQ0YZL6Rn0Yj/jqM5P6LTGThbEQbWQvea1W0cNm3wCYbGYolGuCTJ74sqg6wi5hPKqIJpwYQAsa/MkBb/wH0Bdf8I649ylyGSMouZJPIivSAvOG02vqUMp/sQXQO02Tr4GbfsIgM1AuACNb8e6a6qHq5shEXSKw5KM5ASPge0YZz2rmG2lU0rKu0/h0MDipVOSg28H5zNlQD749g9nS2MrNepYgcHRb+BPy4KPZccB4IDEakYRyvIc4dmdvhHDw54+MTGrYbvuCtAq3iMAfeKrDxZ2dMz88Tg6C+42ARIiz+YQUqW6ANriXWznzmnEawrK+Rzjz9ESdVKft6HpTn76nZnFg0LPthdHnauc8Wqrh0pL5j+e97GJkpwxkNRfJ5aIbc2ohkQYsxUo0kfmkx/5+QmAFYVNTyZ7KRNwiEdF3Fpcwn02a65cFEtkNRjOuC38FgXDoel96FyK8E7CCz3bzNzHaGhjap0KK3E0lmmXYF6KULMf/846erS8n3bMMYc1pdJvvBHBftae1RvPQmL1LMndZuhaWYrQWI4KFAAbYU0b2iFVB07BJzckbEJYrngMTZiDu3UFagLb1TNvJCUIy/RCiRtjzM5XVPlvYN5AfxvOLZyXVUdu44Tn42KNpnQwOKMw6gMfwUlF1UGiI5mxbO7GE/3Ipx5b2j2E/hXC5j9HmWXc6geEVbskN2MmHl9e2FhahqFEYV5LjWqsAgtl3D6WjFbdk2+fa+Xx/uRH2Vjg5uC3mWrVardGGr1NEi6/1CBguvL67SUbqMRrdNh2TC/Xzc3fYRbFjBYoGUKpe1JhnXSUXNJpc3XSJybw7IUXqVji6oSK9G7axzIRqwe5He8XLfreOj7bLrj//2/0D/shHfY+Y1KLu3SzueT+QORybbNT5NJDOatZvNDAI+k24aFr9VSNxd00TWQApmXJDJtEm2JOPWeOUlLq+LAj2ToAZddfw6avxmv/se7sdPMpHQd8gHJ1/bfd4fGH6rsus98GNOdzHwZ5OccdlsOsY3zc6+U5312DVSZ83lnDZN8yeuQmDz","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get Model Status","permalink":"/docs/1backend/get-model-status"},"next":{"title":"List Platforms","permalink":"/docs/1backend/list-platforms"}}');var r=s(74848),n=s(28453),a=s(53746),o=s.n(a),d=s(56518),p=s.n(d),l=s(99972),c=s.n(l),m=s(25342),b=s.n(m),u=(s(44215),s(82223),s(24861));const j={id:"list-models",title:"List Models",description:"Retrieves a list of models.",sidebar_label:"List Models",hide_title:!0,hide_table_of_contents:!0,api:"eJy9Vttu40YM/ZXBPCuSs+kChZ6aFEURbNoEcdIXx8DSEm3PZm7hjOR4Df17QUl2fJFbtCj6Yo95OUNyDklvZImhIOWjclbm8hEjKawxCBBahSjcXBhXog7pi32xj/hWKcIgvrbCi1AXeXvKa4Wrr8IjGRWCcjaViXQeCRj4tpS5ZLjfWiiZSMLgnQ0YZL6Rn0Yj/jqM5P6LTGThbEQbWQvea1W0cNm3wCYbGYolGuCTJ74sqg6wi5hPKqIJpwYQAsa/MkBb/wH0Bdf8I649ylyGSMouZJPIivSAvOG02vqUMp/sQXQO02Tr4GbfsIgM1AuACNb8e6a6qHq5shEXSKw5KM5ASPge0YZz2rmG2lU0rKu0/h0MDipVOSg28H5zNlQD749g9nS2MrNepYgcHRb+BPy4KPZccB4IDEakYRyvIc4dmdvhHDw54+MTGrYbvuCtAq3iMAfeKrDxZ2dMz88Tg6C+42ARIiz+YQUqW6ANriXWznzmnEawrK+Rzjz9ESdVKft6HpTn76nZnFg0LPthdHnauc8Wqrh0pL5j+e97GJkpwxkNRfJ5aIbc2ohkQYsxUo0kfmkx/5+QmAFYVNTyZ7KRNwiEdF3Fpcwn02a65cFEtkNRjOuC38FgXDoel96FyK8E7CCz3bzNzHaGhjap0KK3E0lmmXYF6KULMf/846erS8n3bMMYc1pdJvvBHBftae1RvPQmL1LMndZuhaWYrQWI4KFAAbYU0b2iFVB07BJzckbEJYrngMTZiDu3UFagLb1TNvJCUIy/RCiRtjzM5XVPlvYN5AfxvOLZyXVUdu44Tn42KNpnQwOKMw6gMfwUlF1UGiI5mxbO7GE/3Ipx5b2j2E/hXC5j9HmWXc6geEVbskN2MmHl9e2FhahqFEYV5LjWqsAgtl3D6WjFbdk2+fa+Xx/uRH2Vjg5uC3mWrVardGGr1NEi6/1CBguvL67SUbqMRrdNh2TC/Xzc3fYRbFjBYoGUKpe1JhnXSUXNJpc3XSJybw7IUXqVji6oSK9G7axzIRqwe5He8XLfreOj7bLrj//2/0D/shHfY+Y1KLu3SzueT+QORybbNT5NJDOatZvNDAI+k24aFr9VSNxd00TWQApmXJDJtEm2JOPWeOUlLq+LAj2ToAZddfw6avxmv/se7sdPMpHQd8gHJ1/bfd4fGH6rsus98GNOdzHwZ5OccdlsOsY3zc6+U5312DVSZ83lnDZN8yeuQmDz",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},h=void 0,g={},y=[];function f(e){const t={code:"code",p:"p",...(0,n.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"List Models"}),"\n",(0,r.jsx)(o(),{method:"post",path:"/model-svc/models",context:"endpoint"}),"\n",(0,r.jsx)(t.p,{children:"Retrieves a list of models."}),"\n",(0,r.jsxs)(t.p,{children:["Requires ",(0,r.jsx)(t.code,{children:"model-svc:model:view"})," permission."]}),"\n",(0,r.jsx)(p(),{parameters:void 0}),"\n",(0,r.jsx)(c(),{title:"Body",body:void 0}),"\n",(0,r.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{models:{items:{properties:{assets:{items:{properties:{envVarKey:{type:"string"},url:{type:"string"}},required:["envVarKey","url"],type:"object"},type:"array"},bits:{type:"integer"},description:{type:"string"},extension:{type:"string"},flavour:{type:"string"},fullName:{type:"string"},id:{type:"string"},maxBits:{type:"integer"},maxRam:{type:"number"},mirrors:{items:{type:"string"},type:"array"},name:{type:"string"},parameters:{type:"string"},platformId:{type:"string"},promptTemplate:{type:"string"},quality:{type:"string"},quantComment:{type:"string"},size:{type:"number"},tags:{items:{type:"string"},type:"array"},uncensored:{type:"boolean"},version:{type:"string"}},required:["id","name","platformId"],type:"object"},type:"array"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function v(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(f,{...e})}):f(e)}}}]);