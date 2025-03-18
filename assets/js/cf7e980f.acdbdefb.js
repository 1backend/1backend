"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2641],{52778:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>u,contentTitle:()=>y,default:()=>I,frontMatter:()=>m,metadata:()=>n,toc:()=>g});const n=JSON.parse('{"id":"1backend/register-instance","title":"Register Instance","description":"Registers an instance. Idempotent.","source":"@site/docs/1backend/register-instance.api.mdx","sourceDirName":"1backend","slug":"/1backend/register-instance","permalink":"/docs/1backend/register-instance","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"register-instance","title":"Register Instance","description":"Registers an instance. Idempotent.","sidebar_label":"Register Instance","hide_title":true,"hide_table_of_contents":true,"api":"eJy9Vttu4zYQ/RWCTy3gSE7SAImemmzTrtOgMeL1k2MsaGksc0ORDEnZ6xoC+hH9wn5JMaRkK5bS7rZFEcBQpBmeM7cz3NEMbGq4dlxJmtBHyLl1YCxhknBpHZMpRGSUQaGVA+kiOqBKg2HoMMpoQk3tMqqt6YAaeCnBuhuVbWmyo6mS6IqPTGvBU+8cf7IIuaM2XUHB8EkbPNpxsPhfBlqobQHSIc7uiOmHFZDRD0QtiVsBOdgSt2L4w+2eP+GvwiFqiVHAZ1ZoATTxQB8zuJm+c3fcpHRA3VbjB+sMlzmtBnSlrOtyeK+saxjsT2dZZsDaiDzCS8kNZIQvyfTxHllI5Yg2as0zyF5RKLYWzBpMlKqiD5/7DBzsEe1jxq8Mv7tdvz/rddFdwqPx23QfvBETCXGKNEUliy0Zjb01sAydfSba1C8j/9dHQTO36pIYM7f6exrkG4jyaECeaMw0f6LfvgKNt6o0J3bdWyytTE+xxsr8N8W6HF4O96hcOsjBIKzvY+gCT/z7fwz9uldXzmnbF3RpRBf6x1KIBsiffUSi5+wkjtvNmFxcXZ528aow5EibJjMPPt8bqcUnSB2t0KpfXkgjFj548B11OM+ZEjyA1UraIAZnw9NueO8MMOeL88US06FYDeh3w2HPqMg1Ezwjd5OHX74G4LWGgTHKtHBb+etn0hPlVLLSrZThv35dqP+KyUV/ThwYnM2J7w9y68/8fyjhgEFaGu62NJnt6A0wA+a6RIWZzSvsPpZbbMbQZGZLJusUu7IAt1K4qXSJfRZEicamNkMViflheYXetx7EDxWNY6FSJlD5kovLs/NTinANmzDePqA2p86+2mogT7XJEyVLJYTaQIYKy4jVDCVBZsSpZ5CEpbUmLI0q/MhOLRiMiNyrnEsCMtOKh33M8fwVsAywGJKhCNHrumd8KQ4DzDT/GbY+x1iex8Omvm104Hjz9uzHsA6PtxYuqZ7NhIuotSWa/LcUPAh20NVGRfdKF2rwtjbhqpNL1Vw0WOr7EArG0c8yAfZ7y2VeCuaMkjXXJkvjEZmU2hNoQyVxfLpg6TPIDB1i2pGya1JKvuRYwGBHMliDUNpfQrRgbqlMQZbKEIyUXI9IazbsH7/9XvDUKIyGp2BPFsxC5jtgUXLhcAnblAkv0YKnIK0vTcP7p/E9WZ9Hw06CNptNlMsyUiaPaz8bs1yLk/NoGK1cITAWB6awD8tJQD8EbTcsz8FEXMXeJMbO4c6vh9ObECgdUByQkIZhdB4NT0wanV2F1WtdwWSLaUfy6VEmW/fDL7t+1p3s4LOLtWBcttZfaK4ZbQ+3n5Aae9707ozudpjyqRFVha9fSjCoLPMBXTPD2QJjns3x7hcmC/XgGba4dgLfExxpTAYTZRitI+mrBo3HdZqCdn9p2xaq8fQDHdBFfX0uVIYuhm1wS7INTai/hPs+QgP/bkcFk3nJcrQNR+KQs1qMDuOPjAbNAwbVfJLbFsFj+Qhx4C9G1euy2wVxqaq9ffj0psdes4I1VnJeVdWfLJFRHg==","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"List Definitions","permalink":"/docs/1backend/list-definitions"},"next":{"title":"Remove Instance","permalink":"/docs/1backend/remove-instance"}}');var i=s(74848),a=s(28453),r=s(53746),d=s.n(r),p=s(56518),c=s.n(p),o=s(99972),l=s.n(o),b=s(25342),h=s.n(b),f=(s(44215),s(82223),s(24861));const m={id:"register-instance",title:"Register Instance",description:"Registers an instance. Idempotent.",sidebar_label:"Register Instance",hide_title:!0,hide_table_of_contents:!0,api:"eJy9Vttu4zYQ/RWCTy3gSE7SAImemmzTrtOgMeL1k2MsaGksc0ORDEnZ6xoC+hH9wn5JMaRkK5bS7rZFEcBQpBmeM7cz3NEMbGq4dlxJmtBHyLl1YCxhknBpHZMpRGSUQaGVA+kiOqBKg2HoMMpoQk3tMqqt6YAaeCnBuhuVbWmyo6mS6IqPTGvBU+8cf7IIuaM2XUHB8EkbPNpxsPhfBlqobQHSIc7uiOmHFZDRD0QtiVsBOdgSt2L4w+2eP+GvwiFqiVHAZ1ZoATTxQB8zuJm+c3fcpHRA3VbjB+sMlzmtBnSlrOtyeK+saxjsT2dZZsDaiDzCS8kNZIQvyfTxHllI5Yg2as0zyF5RKLYWzBpMlKqiD5/7DBzsEe1jxq8Mv7tdvz/rddFdwqPx23QfvBETCXGKNEUliy0Zjb01sAydfSba1C8j/9dHQTO36pIYM7f6exrkG4jyaECeaMw0f6LfvgKNt6o0J3bdWyytTE+xxsr8N8W6HF4O96hcOsjBIKzvY+gCT/z7fwz9uldXzmnbF3RpRBf6x1KIBsiffUSi5+wkjtvNmFxcXZ528aow5EibJjMPPt8bqcUnSB2t0KpfXkgjFj548B11OM+ZEjyA1UraIAZnw9NueO8MMOeL88US06FYDeh3w2HPqMg1Ezwjd5OHX74G4LWGgTHKtHBb+etn0hPlVLLSrZThv35dqP+KyUV/ThwYnM2J7w9y68/8fyjhgEFaGu62NJnt6A0wA+a6RIWZzSvsPpZbbMbQZGZLJusUu7IAt1K4qXSJfRZEicamNkMViflheYXetx7EDxWNY6FSJlD5kovLs/NTinANmzDePqA2p86+2mogT7XJEyVLJYTaQIYKy4jVDCVBZsSpZ5CEpbUmLI0q/MhOLRiMiNyrnEsCMtOKh33M8fwVsAywGJKhCNHrumd8KQ4DzDT/GbY+x1iex8Omvm104Hjz9uzHsA6PtxYuqZ7NhIuotSWa/LcUPAh20NVGRfdKF2rwtjbhqpNL1Vw0WOr7EArG0c8yAfZ7y2VeCuaMkjXXJkvjEZmU2hNoQyVxfLpg6TPIDB1i2pGya1JKvuRYwGBHMliDUNpfQrRgbqlMQZbKEIyUXI9IazbsH7/9XvDUKIyGp2BPFsxC5jtgUXLhcAnblAkv0YKnIK0vTcP7p/E9WZ9Hw06CNptNlMsyUiaPaz8bs1yLk/NoGK1cITAWB6awD8tJQD8EbTcsz8FEXMXeJMbO4c6vh9ObECgdUByQkIZhdB4NT0wanV2F1WtdwWSLaUfy6VEmW/fDL7t+1p3s4LOLtWBcttZfaK4ZbQ+3n5Aae9707ozudpjyqRFVha9fSjCoLPMBXTPD2QJjns3x7hcmC/XgGba4dgLfExxpTAYTZRitI+mrBo3HdZqCdn9p2xaq8fQDHdBFfX0uVIYuhm1wS7INTai/hPs+QgP/bkcFk3nJcrQNR+KQs1qMDuOPjAbNAwbVfJLbFsFj+Qhx4C9G1euy2wVxqaq9ffj0psdes4I1VnJeVdWfLJFRHg==",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},y=void 0,u={},g=[];function x(e){const t={p:"p",...(0,a.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"Register Instance"}),"\n",(0,i.jsx)(d(),{method:"put",path:"/registry-svc/instance",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Registers an instance. Idempotent."}),"\n",(0,i.jsx)(f.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(c(),{parameters:void 0}),"\n",(0,i.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{deploymentId:{description:"The ID of the deployment that this instance is an instance of.",example:"depl_deBUCtJirc",type:"string"},host:{description:"Host of the instance address. Required if URL is not provided",example:"myserver.com",type:"string"},id:{example:"inst_di9riJEvH2",type:"string"},ip:{description:"IP of the instance address. Optional: to register by IP instead of host",example:"8.8.8.8",type:"string"},path:{description:'Path of the instance address. Optional (e.g., "/api")',example:"/your-svc",type:"string"},port:{description:"Port of the instance address. Required if URL is not provided",example:8080,type:"integer"},scheme:{description:"Scheme of the instance address. Required if URL is not provided.",example:"https",type:"string"},url:{description:"Full address URL of the instance.",example:"https://myserver.com:5981",type:"string"}},required:["url"],type:"object"}}},description:"Register Instance Request",required:!0}}),"\n",(0,i.jsx)(h(),{id:void 0,label:void 0,responses:{201:{description:"Created",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function I(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(x,{...e})}):x(e)}}}]);