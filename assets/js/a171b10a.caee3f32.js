"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9707],{22972:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>b,contentTitle:()=>g,default:()=>G,frontMatter:()=>k,metadata:()=>i,toc:()=>f});const i=JSON.parse('{"id":"1backend/list-models","title":"List Models","description":"Retrieves a list of models.","source":"@site/docs/1backend/list-models.api.mdx","sourceDirName":"1backend","slug":"/1backend/list-models","permalink":"/docs/1backend/list-models","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"list-models","title":"List Models","description":"Retrieves a list of models.","sidebar_label":"List Models","hide_title":true,"hide_table_of_contents":true,"api":"eJy9Vttu40YM/ZXBPCuSk3SBVk9NiqIINm2COOmLY2BpibZnM7dwRnK8hv59QUl2HFtu0aLoiz3m5QzJOSS9kSWGgpSPylmZyweMpLDGIEBoFaJwc2FciTqkz/bZPuBrpQiD+NIKz0Jd5O0prxWuvgiPZFQIytlUJtJ5JGDgm1LmkuF+b6FkIgmDdzZgkPlGXoxG/PUxkrvPMpGFsxFtZC14r1XRwmVfA5tsZCiWaIBPnviyqDrALmI+qYgmHBtACBj/ygBt/SfQZ1zzj7j2KHMZIim7kE0iK9ID8obTautTynyyB9E5TJOtg5t9xSIyUC8AIljz75nqourlykZcILHmQ3EGQsK3iDac0s411K6iYV2l9R9gcFCpykGxgbfrk6EaeHsAs6ezlZn1KkXk6GPhj8APi2JPBeeBwGBEGsbxGuLckbkZzsGTMz4+omG74QteK9AqDnPgtQIbf3HG9Pw8MgjqGw4WIcLiH1agsgXa4Fpi7cxnzmkEy/oa6cTTH3BSlbKv54fy/D01myOLhmU/jM6PO/fJQhWXjtQ3LP99DyMzZTijoUg+Dc2QGxuRLGgxRqqRxK8t5v8TEjMAi4pa/kw28hqBkK6quJT5ZNpMtzyYyHYoinFd8DsYjEvH49K7EPmVgB1ktpu3mdnO0NAmFVr0diLJLNOuAL10Ieaffry4PJd8zzaMMafVZbIfzGHRHtcexXNv8izF3GntVliK2VqACB4KFGBLEd0LWgFFxy4xJ2dEXKJ4Ckicjbh1C2UF2tI7ZSMvBMX4S4QSacvDXF71ZGnfQL4TzyuenVxHZeeO4+Rng6J9NjSgOOMAGsPPQdlFpSGSs2nhzB72/Y0YV947iv0UzuUyRp9n2fkMihe0JTtkRxNWXt2cWYiqRmFUQY5rrQoMYts1nI5W3JZtk2/v++3+VtSX6ejDbSHPstVqlS5slTpaZL1fyGDh9dllOkqX0ei26ZBMuJuPu9vegw0rWCyQUuWy1iTjOqmo2eT8uktE7s0BOUov09EZFenFT+2scyEasHuR3vJy363jg+2y64//9v9A/7IR32LmNSi7t0s7nk/kDkcm2zU+TSQzmrWbzQwCPpFuGha/VkjcXdNE1kAKZlyQybRJtiTj1njhJS6vigI9k6AGXXX8Omj8Zr/77u/GjzKR0HfIOydf2n3eHxh+q7LrPfBDTncx8GeTnHDZbDrGN83OvlOd9Ng1UmfN5Zw2TfMdvclg+w==","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get Model Status","permalink":"/docs/1backend/get-model-status"},"next":{"title":"List Platforms","permalink":"/docs/1backend/list-platforms"}}');var n=s(74848),r=s(28453),o=s(53746),a=s.n(o),l=s(56518),d=s.n(l),p=s(99972),c=s.n(p),m=s(25342),y=s.n(m),u=(s(44215),s(82223),s(24861));const k={id:"list-models",title:"List Models",description:"Retrieves a list of models.",sidebar_label:"List Models",hide_title:!0,hide_table_of_contents:!0,api:"eJy9Vttu40YM/ZXBPCuSk3SBVk9NiqIINm2COOmLY2BpibZnM7dwRnK8hv59QUl2HFtu0aLoiz3m5QzJOSS9kSWGgpSPylmZyweMpLDGIEBoFaJwc2FciTqkz/bZPuBrpQiD+NIKz0Jd5O0prxWuvgiPZFQIytlUJtJ5JGDgm1LmkuF+b6FkIgmDdzZgkPlGXoxG/PUxkrvPMpGFsxFtZC14r1XRwmVfA5tsZCiWaIBPnviyqDrALmI+qYgmHBtACBj/ygBt/SfQZ1zzj7j2KHMZIim7kE0iK9ID8obTautTynyyB9E5TJOtg5t9xSIyUC8AIljz75nqourlykZcILHmQ3EGQsK3iDac0s411K6iYV2l9R9gcFCpykGxgbfrk6EaeHsAs6ezlZn1KkXk6GPhj8APi2JPBeeBwGBEGsbxGuLckbkZzsGTMz4+omG74QteK9AqDnPgtQIbf3HG9Pw8MgjqGw4WIcLiH1agsgXa4Fpi7cxnzmkEy/oa6cTTH3BSlbKv54fy/D01myOLhmU/jM6PO/fJQhWXjtQ3LP99DyMzZTijoUg+Dc2QGxuRLGgxRqqRxK8t5v8TEjMAi4pa/kw28hqBkK6quJT5ZNpMtzyYyHYoinFd8DsYjEvH49K7EPmVgB1ktpu3mdnO0NAmFVr0diLJLNOuAL10Ieaffry4PJd8zzaMMafVZbIfzGHRHtcexXNv8izF3GntVliK2VqACB4KFGBLEd0LWgFFxy4xJ2dEXKJ4Ckicjbh1C2UF2tI7ZSMvBMX4S4QSacvDXF71ZGnfQL4TzyuenVxHZeeO4+Rng6J9NjSgOOMAGsPPQdlFpSGSs2nhzB72/Y0YV947iv0UzuUyRp9n2fkMihe0JTtkRxNWXt2cWYiqRmFUQY5rrQoMYts1nI5W3JZtk2/v++3+VtSX6ejDbSHPstVqlS5slTpaZL1fyGDh9dllOkqX0ei26ZBMuJuPu9vegw0rWCyQUuWy1iTjOqmo2eT8uktE7s0BOUov09EZFenFT+2scyEasHuR3vJy363jg+2y64//9v9A/7IR32LmNSi7t0s7nk/kDkcm2zU+TSQzmrWbzQwCPpFuGha/VkjcXdNE1kAKZlyQybRJtiTj1njhJS6vigI9k6AGXXX8Omj8Zr/77u/GjzKR0HfIOydf2n3eHxh+q7LrPfBDTncx8GeTnHDZbDrGN83OvlOd9Ng1UmfN5Zw2TfMdvclg+w==",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},g=void 0,b={},f=[];function j(e){const t={code:"code",p:"p",...(0,r.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"List Models"}),"\n",(0,n.jsx)(a(),{method:"post",path:"/model-svc/models",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Retrieves a list of models."}),"\n",(0,n.jsxs)(t.p,{children:["Requires ",(0,n.jsx)(t.code,{children:"model-svc:model:view"})," permission."]}),"\n",(0,n.jsx)(d(),{parameters:void 0}),"\n",(0,n.jsx)(c(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(y(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{models:{items:{properties:{assets:{items:{properties:{envVarKey:{type:"string"},url:{type:"string"}},required:["envVarKey","url"],type:"object"},type:"array"},bits:{type:"integer"},description:{type:"string"},extension:{type:"string"},flavour:{type:"string"},fullName:{type:"string"},id:{type:"string"},maxBits:{type:"integer"},maxRam:{type:"number"},mirrors:{items:{type:"string"},type:"array"},name:{type:"string"},parameters:{type:"string"},platformId:{type:"string"},promptTemplate:{type:"string"},quality:{type:"string"},quantComment:{type:"string"},size:{type:"number"},tags:{items:{type:"string"},type:"array"},uncensored:{type:"boolean"},version:{type:"string"}},required:["id","name","platformId"],type:"object"},type:"array"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function G(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(j,{...e})}):j(e)}}}]);