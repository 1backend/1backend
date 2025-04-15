"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6022],{13627:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>j,contentTitle:()=>L,default:()=>g,frontMatter:()=>h,metadata:()=>n,toc:()=>m});const n=JSON.parse('{"id":"1backend/get-public-key","title":"Get Public Key","description":"Get the public key to parse and verify the JWT.","source":"@site/docs/1backend/get-public-key.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-public-key","permalink":"/docs/1backend/get-public-key","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-public-key","title":"Get Public Key","description":"Get the public key to parse and verify the JWT.","sidebar_label":"Get Public Key","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VE1v2zAM/SsCz46dLAMG+LQWGIq0xVogLXZIc1BkxlYjSyolO8sM//eBdtKv9dTDThbkR/Lx8VEdFBgUaR+1s5DDBUYRKxS+2RitxA4PIjrhJQUU0haiRdLbwwC5/HWXQgLOI0mOXhSQQ4nxdgi9wgMkQBi8swED5B18mU7587bgzRUkoJyNaCP/ld4brYaE2WNgSAdBVVhLPnniclGPCf1zpbyDePAIOYRI2pbQ91z8qdGEBeSrV9B1coK6zSOqCH3P4K8fkVvYVhpdiMvlzU/hSNQ6BG1L4ZGGo7NCF5/nj0SOPub+McXZvxTvrWxi5Uj/wf/FhC9lGVjW+4Aklq1iVWuMlTt6ABLwMlaQQ9YEpEloVTaOYLIbjBGQWiTO0UFDhoGZcUqayoWYz2bz+Tfo14xTDel4WDLxkes5SkI6azj9ezXuDh7FwxHyAGLrjHF7LMTmIKQIXqrRxtHt0AqpRoOILbl68PSpH3HtSm0F2sI7bSP7XHP+CmWBBAlYWbMqZ0ftB5XhWSzpNVuNldJ265gnD0aqYTBYS80dB2kwfGdDNUZGcjZVrn6V+3Yhlo33jljOUaQqRp9n2Wwj1Q5twQEZ9Mk7Fc4WEyujblHUWpFjrbXCILyRceuo5naMVmgDMp9TvYvba9HO0+mbaiHPsv1+n5a2SR2V2TEuZLL0ZjJPp2kVa8McIlIdbrbLsdoL2bCXZYmUapcNkIx10tEwZHY+NgIJsB1G+tN0nk4npFL2QALehVhL+4opP1LjKyPGZ+ZN+93LEnziOTtOMOLvmHkjtWUKgxzd0dIrOFmaXf5i6nUCbF4GdN1GBrwn0/d8/dQgHSBfrRNoJWm54d5X6z45+Ym3gFPkcKYUep53K00zWundFvevV+3ixx30/V/kefla","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Save Permits","permalink":"/docs/1backend/save-permits"},"next":{"title":"Register","permalink":"/docs/1backend/register"}}');var a=i(74848),d=i(28453),s=i(53746),o=i.n(s),r=i(56518),c=i.n(r),l=i(99972),p=i.n(l),b=i(25342),u=i.n(b),k=(i(44215),i(82223),i(24861));const h={id:"get-public-key",title:"Get Public Key",description:"Get the public key to parse and verify the JWT.",sidebar_label:"Get Public Key",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VE1v2zAM/SsCz46dLAMG+LQWGIq0xVogLXZIc1BkxlYjSyolO8sM//eBdtKv9dTDThbkR/Lx8VEdFBgUaR+1s5DDBUYRKxS+2RitxA4PIjrhJQUU0haiRdLbwwC5/HWXQgLOI0mOXhSQQ4nxdgi9wgMkQBi8swED5B18mU7587bgzRUkoJyNaCP/ld4brYaE2WNgSAdBVVhLPnniclGPCf1zpbyDePAIOYRI2pbQ91z8qdGEBeSrV9B1coK6zSOqCH3P4K8fkVvYVhpdiMvlzU/hSNQ6BG1L4ZGGo7NCF5/nj0SOPub+McXZvxTvrWxi5Uj/wf/FhC9lGVjW+4Aklq1iVWuMlTt6ABLwMlaQQ9YEpEloVTaOYLIbjBGQWiTO0UFDhoGZcUqayoWYz2bz+Tfo14xTDel4WDLxkes5SkI6azj9ezXuDh7FwxHyAGLrjHF7LMTmIKQIXqrRxtHt0AqpRoOILbl68PSpH3HtSm0F2sI7bSP7XHP+CmWBBAlYWbMqZ0ftB5XhWSzpNVuNldJ265gnD0aqYTBYS80dB2kwfGdDNUZGcjZVrn6V+3Yhlo33jljOUaQqRp9n2Wwj1Q5twQEZ9Mk7Fc4WEyujblHUWpFjrbXCILyRceuo5naMVmgDMp9TvYvba9HO0+mbaiHPsv1+n5a2SR2V2TEuZLL0ZjJPp2kVa8McIlIdbrbLsdoL2bCXZYmUapcNkIx10tEwZHY+NgIJsB1G+tN0nk4npFL2QALehVhL+4opP1LjKyPGZ+ZN+93LEnziOTtOMOLvmHkjtWUKgxzd0dIrOFmaXf5i6nUCbF4GdN1GBrwn0/d8/dQgHSBfrRNoJWm54d5X6z45+Ym3gFPkcKYUep53K00zWundFvevV+3ixx30/V/kefla",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},L=void 0,j={},m=[];function y(e){const t={p:"p",...(0,d.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(k.default,{as:"h1",className:"openapi__heading",children:"Get Public Key"}),"\n",(0,a.jsx)(o(),{method:"get",path:"/user-svc/public-key",context:"endpoint"}),"\n",(0,a.jsx)(t.p,{children:"Get the public key to parse and verify the JWT."}),"\n",(0,a.jsx)(c(),{parameters:void 0}),"\n",(0,a.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,a.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{publicKey:{type:"string"}},required:["publicKey"],type:"object"}}}},400:{description:"Invalid JSON or missing permission id",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function g(e={}){const{wrapper:t}={...(0,d.R)(),...e.components};return t?(0,a.jsx)(t,{...e,children:(0,a.jsx)(y,{...e})}):y(e)}}}]);