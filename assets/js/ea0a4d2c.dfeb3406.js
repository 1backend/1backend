"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9405],{69336:(t,e,s)=>{s.r(e),s.d(e,{assets:()=>f,contentTitle:()=>m,default:()=>v,frontMatter:()=>u,metadata:()=>o,toc:()=>g});const o=JSON.parse('{"id":"1backend/list-downloads","title":"List Downloads","description":"List download details.","source":"@site/docs/1backend/list-downloads.api.mdx","sourceDirName":"1backend","slug":"/1backend/list-downloads","permalink":"/docs/1backend/list-downloads","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"list-downloads","title":"List Downloads","description":"List download details.","sidebar_label":"List Downloads","hide_title":true,"hide_table_of_contents":true,"api":"eJytVdtu4zYQ/RWCz4pk19sL9NQE2xZBg25Qb54cAzumxjI3FMklR/Y6hv69GMoXOXGKLrYvvnBuZw7PDHeywqiC9qSdlaW805FE5TbWOKhEhQTaxPzRPtq/8UurA0ZBKxSfltrgVVyr8uBbrjVuPgmPodExamdzmUnnMQBnvq1kKY2O9H7vHmUmA0bvbMQoy538YTTirwto3PIIiKOUs4SW2Bm8N1qlAsXnyBE7GdUKG+BfPnB50n3+U4pyJzVhE1/7qIBAWF2n5LT1KEsZKWhbyy47ZsDqZkv7pGdo3587CPyqI0VBTsSV2yTejsz64OqAMQptBWFoIrfJHrZtFhj43yIlARMQqq04VWdily40QLKU2tJP72R2QKstYY2B4WIILlxshO/uL2jwTeM90OpN41Q/4+vef99bhO4FsmyNEUttwQyQC44XUT/jf+xBVxdhHMgbGHva2BgJqI0X41pf/cv9tsFcOO+OuNziMyqSpwMIAbaXPDo+ezcav6bpwUJLKxf0M1bfouWXoLjAj5cm5tYSBmZ9imGNQfyWVPBdlZhTVG3QtJXlbCdvEAKG65Y1Mpt380wS1FGWsyQCMV0rOc9kg7RyPPXeRZKZ9ElTsjgsjmI40zGBjSl9ugZZFMYpMCsXqRyPJ5OfJRc64Jgy3H4Eh2hekvFx61E87l0epVg6Y9wGK7HYChDRg0IBthLkntAKUGnBVWIZXJNE/BAxcD/iztXaCrSVd9oSi1dz/hVChcyuTdMkr/d3m7g9CRq8/hN7oWi7dGnTOEug0nVgA5o7jmAw/hq1rVsDFJzNlWsGue9vxbT13gVmsydpReTLohgvQD2hrTigSJvqjIXr2ysLpNcoGq2CY661wii8AeIZ5HaMVmhjmutDvT/u78R6ko/OqsWyKDabTV7bNnehLvZxsYDam6tJPspX1Jg0I7zUPiynfbUT2LiBusaQa1ckl4J50mTYZXzTNyIzyXLo4Y/yST66Ciqf/JJm30VqwA6Qpkdi+K6ctb87Sf9/ftz2t0v4lQpvQNvBEunVPpOHTDIbPEDzTLKu2b7bLSDiQzBdx8dfWgw8ZPNMriFoWDAts3mXHaTGA/KEW74fpdCzFNZg2l5lL8a6G07h/YfpR5lJaM82O3jN2bLDD05/MNntIPlLZfcY+LPL3gjZ7Xrdd93Rvze9GXEcp96bCZ13XfcPQL773A==","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Pause a Download","permalink":"/docs/1backend/pause-download"},"next":{"title":"Serve a Downloaded file","permalink":"/docs/1backend/serve-download"}}');var d=s(74848),i=s(28453),n=s(53746),a=s.n(n),l=s(56518),r=s.n(l),p=s(99972),c=s.n(p),y=s(25342),h=s.n(y),b=(s(44215),s(82223),s(24861));const u={id:"list-downloads",title:"List Downloads",description:"List download details.",sidebar_label:"List Downloads",hide_title:!0,hide_table_of_contents:!0,api:"eJytVdtu4zYQ/RWCz4pk19sL9NQE2xZBg25Qb54cAzumxjI3FMklR/Y6hv69GMoXOXGKLrYvvnBuZw7PDHeywqiC9qSdlaW805FE5TbWOKhEhQTaxPzRPtq/8UurA0ZBKxSfltrgVVyr8uBbrjVuPgmPodExamdzmUnnMQBnvq1kKY2O9H7vHmUmA0bvbMQoy538YTTirwto3PIIiKOUs4SW2Bm8N1qlAsXnyBE7GdUKG+BfPnB50n3+U4pyJzVhE1/7qIBAWF2n5LT1KEsZKWhbyy47ZsDqZkv7pGdo3587CPyqI0VBTsSV2yTejsz64OqAMQptBWFoIrfJHrZtFhj43yIlARMQqq04VWdily40QLKU2tJP72R2QKstYY2B4WIILlxshO/uL2jwTeM90OpN41Q/4+vef99bhO4FsmyNEUttwQyQC44XUT/jf+xBVxdhHMgbGHva2BgJqI0X41pf/cv9tsFcOO+OuNziMyqSpwMIAbaXPDo+ezcav6bpwUJLKxf0M1bfouWXoLjAj5cm5tYSBmZ9imGNQfyWVPBdlZhTVG3QtJXlbCdvEAKG65Y1Mpt380wS1FGWsyQCMV0rOc9kg7RyPPXeRZKZ9ElTsjgsjmI40zGBjSl9ugZZFMYpMCsXqRyPJ5OfJRc64Jgy3H4Eh2hekvFx61E87l0epVg6Y9wGK7HYChDRg0IBthLkntAKUGnBVWIZXJNE/BAxcD/iztXaCrSVd9oSi1dz/hVChcyuTdMkr/d3m7g9CRq8/hN7oWi7dGnTOEug0nVgA5o7jmAw/hq1rVsDFJzNlWsGue9vxbT13gVmsydpReTLohgvQD2hrTigSJvqjIXr2ysLpNcoGq2CY661wii8AeIZ5HaMVmhjmutDvT/u78R6ko/OqsWyKDabTV7bNnehLvZxsYDam6tJPspX1Jg0I7zUPiynfbUT2LiBusaQa1ckl4J50mTYZXzTNyIzyXLo4Y/yST66Ciqf/JJm30VqwA6Qpkdi+K6ctb87Sf9/ftz2t0v4lQpvQNvBEunVPpOHTDIbPEDzTLKu2b7bLSDiQzBdx8dfWgw8ZPNMriFoWDAts3mXHaTGA/KEW74fpdCzFNZg2l5lL8a6G07h/YfpR5lJaM82O3jN2bLDD05/MNntIPlLZfcY+LPL3gjZ7Xrdd93Rvze9GXEcp96bCZ13XfcPQL773A==",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,f={},g=[];function w(t){const e={code:"code",p:"p",...(0,i.R)(),...t.components};return(0,d.jsxs)(d.Fragment,{children:[(0,d.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"List Downloads"}),"\n",(0,d.jsx)(a(),{method:"post",path:"/file-svc/downloads",context:"endpoint"}),"\n",(0,d.jsx)(e.p,{children:"List download details."}),"\n",(0,d.jsxs)(e.p,{children:["Requires the ",(0,d.jsx)(e.code,{children:"file-svc:download:view"})," permission."]}),"\n",(0,d.jsx)(r(),{parameters:void 0}),"\n",(0,d.jsx)(c(),{title:"Body",body:void 0}),"\n",(0,d.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"List of downloads",content:{"application/json":{schema:{properties:{downloads:{items:{properties:{createdAt:{type:"string"},downloadedBytes:{description:"DownloadedBytes exists to show the download progress in terms of the number of bytes already downloaded.",format:"int64",type:"integer"},error:{type:"string"},fileName:{type:"string"},filePath:{type:"string"},fileSize:{description:"FileSize is the full final downloaded file size.",format:"int64",type:"integer"},id:{type:"string"},progress:{type:"number"},status:{type:"string"},updatedAt:{type:"string"},url:{type:"string"}},type:"object"},type:"array"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function v(t={}){const{wrapper:e}={...(0,i.R)(),...t.components};return e?(0,d.jsx)(e,{...t,children:(0,d.jsx)(w,{...t})}):w(t)}}}]);