"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7385],{80677:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>S,contentTitle:()=>f,default:()=>j,frontMatter:()=>k,metadata:()=>d,toc:()=>b});const d=JSON.parse('{"id":"1backend/make-default","title":"Make a Model Default","description":"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.","source":"@site/docs/1backend/make-default.api.mdx","sourceDirName":"1backend","slug":"/1backend/make-default","permalink":"/docs/1backend/make-default","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"make-default","title":"Make a Model Default","description":"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.","sidebar_label":"Make a Model Default","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VcuO2zYU/RXirlqAluw6RQGtOkGKwk3SGdSZ1cSLa+naYkyRDEnZdQUC+Yh+Yb+kuJJfM+Mp0M1sbEm8j3MPzyE7qCiUXrmorIEC5hSDQNHYirTAIGJNoqIVtjoePv7z7W+xq8kI523jONqTCGSi2KlY2zYKFB/7yNk7eSVfBdEGqjKQYB155MazCgpocEPvhlCQ4NBjQ5F8gOKhe4LyWB8kKH53GGuQYLAhLsSrswokePraKk8VFNG3JCGUNTUIRQdx7zg0RK/MGlJacHBw1gQKvP7DeMx/j9vevgcJpTWRTORVdE6rsp8g/xI4pHvewi6/UBkhpZQkvLlWdma2qFUlfpvf/v5/GjjPBEY1ICbvrb82mnwByeQ5knuDbaytV39R9WpIfrzOSSRvUIs5+S158Utf83UgJQmBytaruO+195bQk79pYw3Fw4KlEnHNsjzocL4tYSGhoVhbFrJrBwFzPOS9GkdhWw5PeXeQZ8pZ8KPqpPjQTzrIvfWac3NtS9S1DbGYTKbTn4CbH7HNedZhvEuET5n8tHckPh9CPoNYWa3tjiqx3AsUwWFJAk0lot2QEVgOjhErb5vevfeBPI8oPti1MoJM5awyMTt6ryasyJ/dd3NQUL8xcOIXnXpP+55xZVaWcfJeYtnvJTWoeOKAmsLPQZl1qzF6a7LSNhe172Zi3jpnPRM2kFTH6Io8nyyx3JCpOCGHJJ+wcDMbGYxqS6JRpbfMtSopCKcxrqxveBytSjKBGM+x3693H8R2mo0fdQtFnu92u2xt2sz6dX7ICzmunR5Ns3FWx0Yzhki+Cber+dDtDDbscL0mnymb9yE586Si5pDJ22EQkMByGOCPs2k2Hvkym77hus6G2KC5QPoRN3Q6eM+n6CMSurN7XvecP0gg0p8xdxqV4Rl6PruDTR7gZBOQwzNIKM4n+SOvLCSwJzit65YY6N7rlPjz15Y8m3YhYYte4ZIp5etDBX6uoFihDvQfxHz3x+HO+F5c3DJXJzgq2+x5r1C3/AYSNrS/uIXSIsmjSRjKsHpTluTiRd6z0yxdHil3959AAh4MfrYUF5PHB65+FdNTSw4Q+DfJF1K6bjBsSqf4YenFjNM5MEQzRYuU0r8SO90A","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get a Model","permalink":"/docs/1backend/get-model"},"next":{"title":"Start a Model","permalink":"/docs/1backend/start-model"}}');var o=a(74848),s=a(28453),n=a(53746),r=a.n(n),i=a(56518),l=a.n(i),p=a(99972),c=a.n(p),u=a(25342),h=a.n(u),m=(a(44215),a(82223),a(24861));const k={id:"make-default",title:"Make a Model Default",description:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.",sidebar_label:"Make a Model Default",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VcuO2zYU/RXirlqAluw6RQGtOkGKwk3SGdSZ1cSLa+naYkyRDEnZdQUC+Yh+Yb+kuJJfM+Mp0M1sbEm8j3MPzyE7qCiUXrmorIEC5hSDQNHYirTAIGJNoqIVtjoePv7z7W+xq8kI523jONqTCGSi2KlY2zYKFB/7yNk7eSVfBdEGqjKQYB155MazCgpocEPvhlCQ4NBjQ5F8gOKhe4LyWB8kKH53GGuQYLAhLsSrswokePraKk8VFNG3JCGUNTUIRQdx7zg0RK/MGlJacHBw1gQKvP7DeMx/j9vevgcJpTWRTORVdE6rsp8g/xI4pHvewi6/UBkhpZQkvLlWdma2qFUlfpvf/v5/GjjPBEY1ICbvrb82mnwByeQ5knuDbaytV39R9WpIfrzOSSRvUIs5+S158Utf83UgJQmBytaruO+195bQk79pYw3Fw4KlEnHNsjzocL4tYSGhoVhbFrJrBwFzPOS9GkdhWw5PeXeQZ8pZ8KPqpPjQTzrIvfWac3NtS9S1DbGYTKbTn4CbH7HNedZhvEuET5n8tHckPh9CPoNYWa3tjiqx3AsUwWFJAk0lot2QEVgOjhErb5vevfeBPI8oPti1MoJM5awyMTt6ryasyJ/dd3NQUL8xcOIXnXpP+55xZVaWcfJeYtnvJTWoeOKAmsLPQZl1qzF6a7LSNhe172Zi3jpnPRM2kFTH6Io8nyyx3JCpOCGHJJ+wcDMbGYxqS6JRpbfMtSopCKcxrqxveBytSjKBGM+x3693H8R2mo0fdQtFnu92u2xt2sz6dX7ICzmunR5Ns3FWx0Yzhki+Cber+dDtDDbscL0mnymb9yE586Si5pDJ22EQkMByGOCPs2k2Hvkym77hus6G2KC5QPoRN3Q6eM+n6CMSurN7XvecP0gg0p8xdxqV4Rl6PruDTR7gZBOQwzNIKM4n+SOvLCSwJzit65YY6N7rlPjz15Y8m3YhYYte4ZIp5etDBX6uoFihDvQfxHz3x+HO+F5c3DJXJzgq2+x5r1C3/AYSNrS/uIXSIsmjSRjKsHpTluTiRd6z0yxdHil3959AAh4MfrYUF5PHB65+FdNTSw4Q+DfJF1K6bjBsSqf4YenFjNM5MEQzRYuU0r8SO90A",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,S={},b=[];function M(e){const t={p:"p",...(0,s.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Make a Model Default"}),"\n",(0,o.jsx)(r(),{method:"put",path:"/model-svc/model/{modelId}/make-default",context:"endpoint"}),"\n",(0,o.jsx)(t.p,{children:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used."}),"\n",(0,o.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,o.jsx)(l(),{parameters:[{description:"Model ID",in:"path",name:"modelId",required:!0,schema:{type:"string"}}]}),"\n",(0,o.jsx)(c(),{title:"Body",body:void 0}),"\n",(0,o.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function j(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,o.jsx)(t,{...e,children:(0,o.jsx)(M,{...e})}):M(e)}}}]);