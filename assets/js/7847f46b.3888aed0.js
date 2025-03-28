"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2595],{33767:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>j,contentTitle:()=>m,default:()=>x,frontMatter:()=>g,metadata:()=>a,toc:()=>b});const a=JSON.parse('{"id":"1backend/add-message","title":"Add Message","description":"Add a new message to a specific thread.","source":"@site/docs/1backend/add-message.api.mdx","sourceDirName":"1backend","slug":"/1backend/add-message","permalink":"/docs/1backend/add-message","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"add-message","title":"Add Message","description":"Add a new message to a specific thread.","sidebar_label":"Add Message","hide_title":true,"hide_table_of_contents":true,"api":"eJylVltv2zYU/isEX7YBiuQ0K1DoZXO7dvNuCeoWfUiM4YQ8lthKpEpSdj1D/304pGQrkVYs2Ishk+fyndt3eOQSnbCq8cponvOllAyYxj2r0TkokHnDgLkGhdoqwXxpEWTKE24atEBaK8lzDlL+ERV4whuwUKNH63h+e3zk4V2wwFY/8YQrOmjAlzzhGmrkOY8OVpIn3OLnVlmUPPe2xYQ7UWINPD9yf2hI1nmrdMG7bhOF0fmXRh5IQhjtUXv6hKaplAhQs4+OMBxHphpLgXiFjv71QU8vhEXwKJd+xn3Ct6rClQyCD4N9Ey+YxK3S6JgvkZEwA+9BlDVqHw+HbJfgKLnKY+1mXfUHYC0c6L+SJIZfoG4quqhd8RfW6+ub6sPzF4YnUws1+hA5SKkIJlQ3o1Bjqnslc/8RhQ9u8YufhvcOv3jWp5qZ7YNIsEjZHf9FJWxfgv/Gsbb54Y7PATqVfGq/v3lkmzJ0jtiXNkSMr99/WBZzDtpGfqV4rUM75/x9OGcqFkidQJA825eG7a3x+ADXnX5jLFuuhgNSVY5tFVbBENaNP6RTiN242W+ppqOsbCbl6KYnXTIzyP1EsrdxNiYjFdy6xmgXm/zZYjFNw2DEtUKgc9u2qg4MpEQC+Z/H7AnNFoL5fg7KSu+gUpL9ur7+8ynOH2c7OricKbmG1pfGqr+fFt2cg+fzEXi0Giq2RrtDy15ba+z/89Ql3KForfKHwLYvESzaZetLnt9uiBo9FETE/FUJnq13gjqqRl8aIu7GhMYIJJzzTJTgL9xOZLH9suPQhl1Wn/jdBfCR3FtbkV5WGQFVaZzPn794dnXJyfGAa03wY4uN0U2m/dAgu+tF7jjbmqoye5Ts/hBWEAhkoCXz5hNqBiK2MttaU4cppIGl+NjvplCaoZaNUdqnw6IpESTa86pZ9rUOuT4PJTTqNzyEKaP8vz0vltcD5zxYFKPVMNRltBFuh7NN5OopQfd8fCLZs5EzMU5JbsRpZ4WBysbEovTWDBsRRGgxrEFR1RxU6H50ShdtBd4anQpTj/Jzs2LrtmmMpQ6JhS69b/Isu7wH8Qm1JIWMT8lndaHBqx2yWglrqF+UQMeaCvzW2JpKUimB2oUMDv5+vvmd7a7SxQNvLs+y/X6fFrpNjS2yXs9lUDTVxVW6SEtfV3FH2dpdb9fR2xms20NRoE2VyYJIRqlVPiyPy5cxEJ5waukIf5FepYsLK9KrBdmlEalBj5GeyZU/in309HjCY6pvPWqArKlA6bCYKAfHfjRv+TCap8bgCc9Hz6WhJTcJpzkklePxHhy+t1XX0fHnFi2RxCbhO7AK7ikF9EBTjr4lz7dQOfxKRN++7RfId2z8jpuFP4yTPlByoWrpH0/4JzyM33ndpkuG0SQw8fpVdHlBnDBSn5BjlwwaSyGw8V+VHfPezfX6HU/4ff9erI0kHQt72pKwj1BNSEEgrnB25BXoog2Dz6NNGjHo6exMIAQpGT4oqtlkPCagGAj9UlizKsdjpKeuO8nHq3/VOLFelKbabLqu+wfPJheS","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Update Thread","permalink":"/docs/1backend/update-thread"},"next":{"title":"List Messages","permalink":"/docs/1backend/get-messages"}}');var d=s(74848),n=s(28453),i=s(53746),p=s.n(i),r=s(56518),o=s.n(r),l=s(99972),c=s.n(l),h=s(25342),f=s.n(h),u=(s(44215),s(82223),s(24861));const g={id:"add-message",title:"Add Message",description:"Add a new message to a specific thread.",sidebar_label:"Add Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVltv2zYU/isEX7YBiuQ0K1DoZXO7dvNuCeoWfUiM4YQ8lthKpEpSdj1D/304pGQrkVYs2Ishk+fyndt3eOQSnbCq8cponvOllAyYxj2r0TkokHnDgLkGhdoqwXxpEWTKE24atEBaK8lzDlL+ERV4whuwUKNH63h+e3zk4V2wwFY/8YQrOmjAlzzhGmrkOY8OVpIn3OLnVlmUPPe2xYQ7UWINPD9yf2hI1nmrdMG7bhOF0fmXRh5IQhjtUXv6hKaplAhQs4+OMBxHphpLgXiFjv71QU8vhEXwKJd+xn3Ct6rClQyCD4N9Ey+YxK3S6JgvkZEwA+9BlDVqHw+HbJfgKLnKY+1mXfUHYC0c6L+SJIZfoG4quqhd8RfW6+ub6sPzF4YnUws1+hA5SKkIJlQ3o1Bjqnslc/8RhQ9u8YufhvcOv3jWp5qZ7YNIsEjZHf9FJWxfgv/Gsbb54Y7PATqVfGq/v3lkmzJ0jtiXNkSMr99/WBZzDtpGfqV4rUM75/x9OGcqFkidQJA825eG7a3x+ADXnX5jLFuuhgNSVY5tFVbBENaNP6RTiN242W+ppqOsbCbl6KYnXTIzyP1EsrdxNiYjFdy6xmgXm/zZYjFNw2DEtUKgc9u2qg4MpEQC+Z/H7AnNFoL5fg7KSu+gUpL9ur7+8ynOH2c7OricKbmG1pfGqr+fFt2cg+fzEXi0Giq2RrtDy15ba+z/89Ql3KForfKHwLYvESzaZetLnt9uiBo9FETE/FUJnq13gjqqRl8aIu7GhMYIJJzzTJTgL9xOZLH9suPQhl1Wn/jdBfCR3FtbkV5WGQFVaZzPn794dnXJyfGAa03wY4uN0U2m/dAgu+tF7jjbmqoye5Ts/hBWEAhkoCXz5hNqBiK2MttaU4cppIGl+NjvplCaoZaNUdqnw6IpESTa86pZ9rUOuT4PJTTqNzyEKaP8vz0vltcD5zxYFKPVMNRltBFuh7NN5OopQfd8fCLZs5EzMU5JbsRpZ4WBysbEovTWDBsRRGgxrEFR1RxU6H50ShdtBd4anQpTj/Jzs2LrtmmMpQ6JhS69b/Isu7wH8Qm1JIWMT8lndaHBqx2yWglrqF+UQMeaCvzW2JpKUimB2oUMDv5+vvmd7a7SxQNvLs+y/X6fFrpNjS2yXs9lUDTVxVW6SEtfV3FH2dpdb9fR2xms20NRoE2VyYJIRqlVPiyPy5cxEJ5waukIf5FepYsLK9KrBdmlEalBj5GeyZU/in309HjCY6pvPWqArKlA6bCYKAfHfjRv+TCap8bgCc9Hz6WhJTcJpzkklePxHhy+t1XX0fHnFi2RxCbhO7AK7ikF9EBTjr4lz7dQOfxKRN++7RfId2z8jpuFP4yTPlByoWrpH0/4JzyM33ndpkuG0SQw8fpVdHlBnDBSn5BjlwwaSyGw8V+VHfPezfX6HU/4ff9erI0kHQt72pKwj1BNSEEgrnB25BXoog2Dz6NNGjHo6exMIAQpGT4oqtlkPCagGAj9UlizKsdjpKeuO8nHq3/VOLFelKbabLqu+wfPJheS",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,j={},b=[];function y(e){const t={p:"p",...(0,n.R)(),...e.components};return(0,d.jsxs)(d.Fragment,{children:[(0,d.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"Add Message"}),"\n",(0,d.jsx)(p(),{method:"post",path:"/chat-svc/thread/{threadId}/message",context:"endpoint"}),"\n",(0,d.jsx)(t.p,{children:"Add a new message to a specific thread."}),"\n",(0,d.jsx)(u.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,d.jsx)(o(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,d.jsx)(c(),{title:"Body",body:{content:{"application/json":{schema:{properties:{message:{properties:{createdAt:{type:"string"},fileIds:{description:"FileIds defines the file attachments the message has.",items:{type:"string"},type:"array"},id:{example:"msg_emSOPlW58o",type:"string"},meta:{additionalProperties:!0,type:"object"},text:{description:'Text content of the message eg. "Hi, what\'s up?"',type:"string"},threadId:{description:"ThreadId of the message.",example:"thr_emSOeEUWAg",type:"string"},updatedAt:{type:"string"},userId:{description:"UserId is the id of the user who wrote the message.\nFor AI messages this field is empty.",type:"string"}},required:["id","threadId"],type:"object"}},type:"object"}}},description:"Add Message Request",required:!0}}),"\n",(0,d.jsx)(f(),{id:void 0,label:void 0,responses:{200:{description:"Message successfully added",content:{"application/json":{schema:{additionalProperties:!0,type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function x(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,d.jsx)(t,{...e,children:(0,d.jsx)(y,{...e})}):y(e)}}}]);