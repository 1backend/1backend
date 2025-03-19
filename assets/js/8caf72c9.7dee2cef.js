"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2482],{71385:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>R,contentTitle:()=>b,default:()=>k,frontMatter:()=>m,metadata:()=>s,toc:()=>K});const s=JSON.parse('{"id":"1backend/delete-user","title":"Delete a User","description":"Delete a user based on the user ID.","source":"@site/docs/1backend/delete-user.api.mdx","sourceDirName":"1backend","slug":"/1backend/delete-user","permalink":"/docs/1backend/delete-user","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"delete-user","title":"Delete a User","description":"Delete a user based on the user ID.","sidebar_label":"Delete a User","hide_title":true,"hide_table_of_contents":true,"api":"eJylVE1v4zYQ/SvEnFqAkZykC7Q61YsEhbEBdlFvTlkfaGoscUORXJKy6wr874sRpdj52KJoL7YkznDevPdmBqgxSK9cVNZABTeoMSITrA/o2VYErJk1LLaYv6xuCuBgHXpBGasaKqjHnPuAHjg44UWHEX2A6mF4cft9vgI4KHp1IrbAwYgOoQK6f1UDB4/feuWxhir6HjkE2WInoBogHh1FhuiVaSClDQUHZ03AQOdXiwX9PS/68QNwkNZENJFOhXNayRF++TVQyPC6hN1+RRkhpZQ4/LK4fH3tvRF9bK1Xf2P9Hwo89UAF3r2Fe2UieiM0W6Pfo2e33lv//yolDgFl71U8juq8R+HRL/vYQvWwITqjaEi4rNR6L2HDocPY2pPQo8iUASVJdhH2cnwoh6xgAqpCkLMFeq8pttRWCt3aEKt3v15dXwKVm9GsCXTW8BzTS0o+Hx2yL1PIF2A7q7U9YM22RyZYcEIiE6Zm0T6iYUJmH7Gdt91o4bkrdmcbZRia2lllYjEbskVRjy6eLLmcFB4ZBj4zKpz6gEcgOpXZWcJJogg5ioKdUNRxEBrD70GZptciemsKabuzuz+t2Lp3zvoIfCKpjdFVZXm5FfIRTU0JJST+goUl643aKWo7x7Ea96it69BE5rSIO+s7trOedUp6S2IoieEij/Nyxc58E6h5rSSagIR+RvfHpzu2vy4Wz7CFqiwPh0PRmL6wvimnvFCKxumL62JRtLHThDii78LH3TqXPrUWDqJp0BfKlmNISayqqCnk8n1uBziQeXKzi+K6WFx4WVz9Rvc6G2InzBnSp401LaBnXA2nafmXq22SOOJfsXRaKENVRwaGyfYPMNueuMlFq2l5bTiQwyloGKjEvdcp0edvPXoaug2HvfBKbKllWpAq0HMN1U7ogP+A/6c/p734Mzvt0TfhzjY1R6JS6J7egMMjHk97Nm0Snw1PQPLhUkp08Szt1YpJ5xvh5vbu9vMtcBDTvJ4mhO7j8wMVeBPVywnLKOg38R+kDEOev5Se4vPRDzOexjpHE0mblNJ3b8ld6Q==","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Create a New User","permalink":"/docs/1backend/create-user"},"next":{"title":"Save User","permalink":"/docs/1backend/save-user"}}');var n=a(74848),i=a(28453),r=a(53746),o=a.n(r),d=a(56518),l=a.n(d),c=a(99972),p=a.n(c),u=a(25342),v=a.n(u),h=(a(44215),a(82223),a(24861));const m={id:"delete-user",title:"Delete a User",description:"Delete a user based on the user ID.",sidebar_label:"Delete a User",hide_title:!0,hide_table_of_contents:!0,api:"eJylVE1v4zYQ/SvEnFqAkZykC7Q61YsEhbEBdlFvTlkfaGoscUORXJKy6wr874sRpdj52KJoL7YkznDevPdmBqgxSK9cVNZABTeoMSITrA/o2VYErJk1LLaYv6xuCuBgHXpBGasaKqjHnPuAHjg44UWHEX2A6mF4cft9vgI4KHp1IrbAwYgOoQK6f1UDB4/feuWxhir6HjkE2WInoBogHh1FhuiVaSClDQUHZ03AQOdXiwX9PS/68QNwkNZENJFOhXNayRF++TVQyPC6hN1+RRkhpZQ4/LK4fH3tvRF9bK1Xf2P9Hwo89UAF3r2Fe2UieiM0W6Pfo2e33lv//yolDgFl71U8juq8R+HRL/vYQvWwITqjaEi4rNR6L2HDocPY2pPQo8iUASVJdhH2cnwoh6xgAqpCkLMFeq8pttRWCt3aEKt3v15dXwKVm9GsCXTW8BzTS0o+Hx2yL1PIF2A7q7U9YM22RyZYcEIiE6Zm0T6iYUJmH7Gdt91o4bkrdmcbZRia2lllYjEbskVRjy6eLLmcFB4ZBj4zKpz6gEcgOpXZWcJJogg5ioKdUNRxEBrD70GZptciemsKabuzuz+t2Lp3zvoIfCKpjdFVZXm5FfIRTU0JJST+goUl643aKWo7x7Ea96it69BE5rSIO+s7trOedUp6S2IoieEij/Nyxc58E6h5rSSagIR+RvfHpzu2vy4Wz7CFqiwPh0PRmL6wvimnvFCKxumL62JRtLHThDii78LH3TqXPrUWDqJp0BfKlmNISayqqCnk8n1uBziQeXKzi+K6WFx4WVz9Rvc6G2InzBnSp401LaBnXA2nafmXq22SOOJfsXRaKENVRwaGyfYPMNueuMlFq2l5bTiQwyloGKjEvdcp0edvPXoaug2HvfBKbKllWpAq0HMN1U7ogP+A/6c/p734Mzvt0TfhzjY1R6JS6J7egMMjHk97Nm0Snw1PQPLhUkp08Szt1YpJ5xvh5vbu9vMtcBDTvJ4mhO7j8wMVeBPVywnLKOg38R+kDEOev5Se4vPRDzOexjpHE0mblNJ3b8ld6Q==",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},b=void 0,R={},K=[];function f(e){const t={p:"p",...(0,i.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Delete a User"}),"\n",(0,n.jsx)(o(),{method:"delete",path:"/user-svc/user/{userId}",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Delete a user based on the user ID."}),"\n",(0,n.jsx)(h.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(l(),{parameters:[{description:"User ID",in:"path",name:"userId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(v(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function k(e={}){const{wrapper:t}={...(0,i.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(f,{...e})}):f(e)}}}]);