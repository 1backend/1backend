"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5424],{76373:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>y,contentTitle:()=>g,default:()=>V,frontMatter:()=>b,metadata:()=>s,toc:()=>I});const s=JSON.parse('{"id":"1backend/get-thread","title":"Get Thread","description":"Fetch information about a specific chat thread by its ID","source":"@site/docs/1backend/get-thread.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-thread","permalink":"/docs/1backend/get-thread","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-thread","title":"Get Thread","description":"Fetch information about a specific chat thread by its ID","sidebar_label":"Get Thread","hide_title":true,"hide_table_of_contents":true,"api":"eJylVt9v2zYQ/lcIPm2AIjnNCgx6mrt1gbti7eb0KTGGM3WW2FAkS57seoL+9+Ek2VEStVjRl8Qm78d3H+++cysLjCpoT9pZmcvfkVQltN25UAOfCdi6hgSI6FHpnVZCVUCCqoBQiO1RaIpi9ZtMpPMYepdVIXNZIt30NjKRHgLUSBiizG/bJxkHqyGE5gMPVMlEWqhR5nJItOIwAT81OmAhcwoNJjKqCmuQeSvp6Nk2UtC2lF23YePonY0Y+f7FYsH/ZhMXSKBNFLFRCmPcNcYcRUAKGvfIaZWzhJY4AHhvtOqLzD5GjtJOUPjAFJAecuJnHSlO0G2dMwhWdslY1HMfFRAIiyXNFJVIXQxxofZmpOYfrNd/Wf9mW13L5LkHaWLLZ5XzsXA7QRWOL5nOujuv1aqIMxHGG1HgTluM4lBpVYneIU7Cii0aZ8soyKV39ma4VmBFcE1ZmaPYsiV/JsYDURCUUexcGANEBqYJ6zhLyXgAIcCRvze++AqBTcQwW86H4aIHrot4oobtuTTXQ47IWHWcMPZ/gXXT5r3lh9ycTdz2IyrqbZ6e8NlPc627snswuhBv1u/+/JYOfTonQ4LLGUIsNFS5oP/9thGYS/ByvgLCYMGINYY9BvE6BBe+L1OXyIiqCZqOvcy8QggYlg1VMr/dsCZwbzH/v7KCrfeKX6FGqtwoWL1UsbnMWOQu4l5lw2Nn7UmHOslpGPMgZk0wbJ8Zp8BULlL+8ucXV5eS853grBn1MOFTUM+G6uhR3I0md1LsnDHugL3Ksv6CQgG2EOTu0QpQQ0OJXXB1367cxVyWeOtKbQXawjttKT0Ja4VQYHiQ1uX4xD3FD/MPXv+BQ9PyHuh1yVkC1b8K1qC54ggG4y9R27IxQMHZVLl6Evv9Sqwb711gVgeSKiKfZ9nlFtQ92oIdMp6VxywsVxcWSO9R1FoFx1xrhVF4A8RbicsxWqGNvbKd8l2/fyv2V+niUbaYZ9nhcEhL26QulNnoFzMovbm4ShdpRbXp5xVDHd/t1kO2B7DxAGWJIdUu600yeVZVeflqKEQmktthgL9Ir9LFRVDp1YLjehepBjtBeo0kzpvxUentQ/d/zx4en5HwM2XegO43Ts9JO7b3rTy1tzzvokTm51W7SSR3Mhu27RYifgim6/j4U4OBp2uTyD0EDVsmgle6jvy5kPkOTMSvFPbD36MQ/iimm38W9Kkh7ZEpBtPwN5nIezxOfxl0my45NTeDGa6XSqGnieMzPemm43/9+kYmEsbBfBgFDpacPnD0WVBPR2mAwH+75AsubTsMWted7YerL3qc53ewZo42Xdf9BzZoV4A=","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Delete a Thread","permalink":"/docs/1backend/delete-thread"},"next":{"title":"Update Thread","permalink":"/docs/1backend/update-thread"}}');var a=i(74848),n=i(28453),d=i(53746),r=i.n(d),c=i(56518),o=i.n(c),p=i(99972),h=i.n(p),l=i(25342),m=i.n(l),u=(i(44215),i(82223),i(24861));const b={id:"get-thread",title:"Get Thread",description:"Fetch information about a specific chat thread by its ID",sidebar_label:"Get Thread",hide_title:!0,hide_table_of_contents:!0,api:"eJylVt9v2zYQ/lcIPm2AIjnNCgx6mrt1gbti7eb0KTGGM3WW2FAkS57seoL+9+Ek2VEStVjRl8Qm78d3H+++cysLjCpoT9pZmcvfkVQltN25UAOfCdi6hgSI6FHpnVZCVUCCqoBQiO1RaIpi9ZtMpPMYepdVIXNZIt30NjKRHgLUSBiizG/bJxkHqyGE5gMPVMlEWqhR5nJItOIwAT81OmAhcwoNJjKqCmuQeSvp6Nk2UtC2lF23YePonY0Y+f7FYsH/ZhMXSKBNFLFRCmPcNcYcRUAKGvfIaZWzhJY4AHhvtOqLzD5GjtJOUPjAFJAecuJnHSlO0G2dMwhWdslY1HMfFRAIiyXNFJVIXQxxofZmpOYfrNd/Wf9mW13L5LkHaWLLZ5XzsXA7QRWOL5nOujuv1aqIMxHGG1HgTluM4lBpVYneIU7Cii0aZ8soyKV39ma4VmBFcE1ZmaPYsiV/JsYDURCUUexcGANEBqYJ6zhLyXgAIcCRvze++AqBTcQwW86H4aIHrot4oobtuTTXQ47IWHWcMPZ/gXXT5r3lh9ycTdz2IyrqbZ6e8NlPc627snswuhBv1u/+/JYOfTonQ4LLGUIsNFS5oP/9thGYS/ByvgLCYMGINYY9BvE6BBe+L1OXyIiqCZqOvcy8QggYlg1VMr/dsCZwbzH/v7KCrfeKX6FGqtwoWL1UsbnMWOQu4l5lw2Nn7UmHOslpGPMgZk0wbJ8Zp8BULlL+8ucXV5eS853grBn1MOFTUM+G6uhR3I0md1LsnDHugL3Ksv6CQgG2EOTu0QpQQ0OJXXB1367cxVyWeOtKbQXawjttKT0Ja4VQYHiQ1uX4xD3FD/MPXv+BQ9PyHuh1yVkC1b8K1qC54ggG4y9R27IxQMHZVLl6Evv9Sqwb711gVgeSKiKfZ9nlFtQ92oIdMp6VxywsVxcWSO9R1FoFx1xrhVF4A8RbicsxWqGNvbKd8l2/fyv2V+niUbaYZ9nhcEhL26QulNnoFzMovbm4ShdpRbXp5xVDHd/t1kO2B7DxAGWJIdUu600yeVZVeflqKEQmktthgL9Ir9LFRVDp1YLjehepBjtBeo0kzpvxUentQ/d/zx4en5HwM2XegO43Ts9JO7b3rTy1tzzvokTm51W7SSR3Mhu27RYifgim6/j4U4OBp2uTyD0EDVsmgle6jvy5kPkOTMSvFPbD36MQ/iimm38W9Kkh7ZEpBtPwN5nIezxOfxl0my45NTeDGa6XSqGnieMzPemm43/9+kYmEsbBfBgFDpacPnD0WVBPR2mAwH+75AsubTsMWted7YerL3qc53ewZo42Xdf9BzZoV4A=",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},g=void 0,y={},I=[];function f(e){const t={p:"p",...(0,n.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"Get Thread"}),"\n",(0,a.jsx)(r(),{method:"get",path:"/chat-svc/thread/{threadId}",context:"endpoint"}),"\n",(0,a.jsx)(t.p,{children:"Fetch information about a specific chat thread by its ID"}),"\n",(0,a.jsx)(u.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,a.jsx)(o(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,a.jsx)(h(),{title:"Body",body:void 0}),"\n",(0,a.jsx)(m(),{id:void 0,label:void 0,responses:{200:{description:"Thread details successfully retrieved",content:{"application/json":{schema:{properties:{exists:{type:"boolean"},thread:{properties:{createdAt:{type:"string"},id:{example:"thr_emSQnpJbhG",type:"string"},title:{description:"Title of the thread.",type:"string"},topicIds:{description:"TopicIds defines which topics the thread belongs to.\nTopics can roughly be thought of as tags for threads.",items:{type:"string"},type:"array"},updatedAt:{type:"string"},userIds:{description:"UserIds the ids of the users who can see this thread.",items:{type:"string"},type:"array"}},required:["id"],type:"object"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function V(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,a.jsx)(t,{...e,children:(0,a.jsx)(f,{...e})}):f(e)}}}]);