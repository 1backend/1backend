"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5735],{86114:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>b,contentTitle:()=>u,default:()=>G,frontMatter:()=>f,metadata:()=>a,toc:()=>L});const a=JSON.parse('{"id":"1backend/get-messages","title":"List Messages","description":"Fetch messages (and associated assets) for a specific chat thread.","source":"@site/docs/1backend/get-messages.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-messages","permalink":"/docs/1backend/get-messages","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-messages","title":"List Messages","description":"Fetch messages (and associated assets) for a specific chat thread.","sidebar_label":"List Messages","hide_title":true,"hide_table_of_contents":true,"api":"eJytVt+P2zYM/lcEvawFfPb9WIEuL1u6tVu26+6w9NCHu2DgybStVpZUiU6aBf7fB9pO4lzSAsX2ktgSRX78SH7WRuYYVdCetLNyIt8gqUrUGCOUGMUzsLmAGJ3SQNg9IsXnonBBgIgelS60EqoCElQFhDyViXQeA7DDWS4nskR6O/iTifQQoEbCEOXkfvMk+rvOhZj9IhOpecEDVTKRFmqUE9lHmOUykQE/NTpgLicUGkxkVBXWICcbSWvPtpGCtqVs2wUbR+9sxMj7l+fn/HcYeAtQDPkiRREbpTDGojFmLQJS0LhEjq2cJbTEXsB7o1WXa/YhsqvNCIoPzATpPvCWVH7WhHU8NlEBmeYpnUgkkYU2OMvjMfo3/YbIsdAWo6AKBRsLIAJV1WipXxwgiAoi12mH4ijUsAAhwJrfdc5m+Blqb3ijjuXfWM9vbs37Fy+dTI491EgdB5DnmmGCuR2l2hdtOOQeP6CiLix+puP03uFnEgPpwhUHmWCZigf5m07EqgL6LorG//ggTwHaNc+x/2HniW9maJ8xVaHLGF/fvZ+WpwI0Pv9K8ZqI4VTwu25d6L5AegeC7cWqcmIVHOEBrgf7xgUxne3HlCodRaHRdI6w9rROjyG247G555qOWFmcKsdBF7RHFi2vfX9qnmZ2CUbn4vf5zZ/fMjFPEfcBLk7QZqGhygX9z7eN5KkAL05nQBgsGDHHsMQgXofgwn+L1CYyomqCpnWnfa8QAoZpQ5Wc3C9YqAhKlkX5M+vpfKm4KjVS5VhHvYvUCSjby4w19ywuVdaXMNtsS9lm9V5uY4e+19omGD6YGafAVC7S5MXLy6sLyZG3wOaMvxejMbyjkVl7FA+DyYMUhTPGrTAXj+vuuwAKOykl9xGtANV3nSiCq7tW5q7nBMW1K7UVaHPvtKV0q/sVQo5hr/zTodgd2fvOBq//wL41tS1cJ6HOEqiuPliD5owjGIw/RW3LxgAFZ1Pl6pHv25mYN967wPT2JFVEfpJlF4+gPqLN+UDGE3HIwnR2ZoH0EkWtVXDMtVYYhTdAhQs1p2O0QhuR8Wzj/Xp7LZZX6flBtDjJstVqlZa2SV0os+FczKD05uwqPU8rqk0vkqGON8W8j7YHG1dQlhhS7bLOJGOeNHXqdfGqT0Qmktuhh3+eXqXnZ0Gllz+wX+6vGuwI6bWOJEbf7oPsN/tR+H9uDUNN+RuQeQPadrLJBG2Gpr+X26bfKZdM5GR0Ldh1/iKR3OF8ZrN5hIh3wbQtL39qMPD8LRK5hKDhkQnim4iO/JzLSQEm4leyffbXoKLPxfjCchL/tlHtmqkH0/CbTORHXI8vNO2iTbZNz2D67alS6Gl08Ehx2rFA3N7M38lEwjCx+xlhb8n2gd2fRPV0xnoM/NsmXziy2fQT2LY7+37riyd2g91bM0mLtm3/BU1LmrE=","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Add Message","permalink":"/docs/1backend/add-message"},"next":{"title":"Get Threads","permalink":"/docs/1backend/get-threads"}}');var i=t(74848),n=t(28453),d=t(53746),o=t.n(d),r=t(56518),c=t.n(r),p=t(99972),h=t.n(p),l=t(25342),m=t.n(l),g=(t(44215),t(82223),t(24861));const f={id:"get-messages",title:"List Messages",description:"Fetch messages (and associated assets) for a specific chat thread.",sidebar_label:"List Messages",hide_title:!0,hide_table_of_contents:!0,api:"eJytVt+P2zYM/lcEvawFfPb9WIEuL1u6tVu26+6w9NCHu2DgybStVpZUiU6aBf7fB9pO4lzSAsX2ktgSRX78SH7WRuYYVdCetLNyIt8gqUrUGCOUGMUzsLmAGJ3SQNg9IsXnonBBgIgelS60EqoCElQFhDyViXQeA7DDWS4nskR6O/iTifQQoEbCEOXkfvMk+rvOhZj9IhOpecEDVTKRFmqUE9lHmOUykQE/NTpgLicUGkxkVBXWICcbSWvPtpGCtqVs2wUbR+9sxMj7l+fn/HcYeAtQDPkiRREbpTDGojFmLQJS0LhEjq2cJbTEXsB7o1WXa/YhsqvNCIoPzATpPvCWVH7WhHU8NlEBmeYpnUgkkYU2OMvjMfo3/YbIsdAWo6AKBRsLIAJV1WipXxwgiAoi12mH4ijUsAAhwJrfdc5m+Blqb3ijjuXfWM9vbs37Fy+dTI491EgdB5DnmmGCuR2l2hdtOOQeP6CiLix+puP03uFnEgPpwhUHmWCZigf5m07EqgL6LorG//ggTwHaNc+x/2HniW9maJ8xVaHLGF/fvZ+WpwI0Pv9K8ZqI4VTwu25d6L5AegeC7cWqcmIVHOEBrgf7xgUxne3HlCodRaHRdI6w9rROjyG247G555qOWFmcKsdBF7RHFi2vfX9qnmZ2CUbn4vf5zZ/fMjFPEfcBLk7QZqGhygX9z7eN5KkAL05nQBgsGDHHsMQgXofgwn+L1CYyomqCpnWnfa8QAoZpQ5Wc3C9YqAhKlkX5M+vpfKm4KjVS5VhHvYvUCSjby4w19ywuVdaXMNtsS9lm9V5uY4e+19omGD6YGafAVC7S5MXLy6sLyZG3wOaMvxejMbyjkVl7FA+DyYMUhTPGrTAXj+vuuwAKOykl9xGtANV3nSiCq7tW5q7nBMW1K7UVaHPvtKV0q/sVQo5hr/zTodgd2fvOBq//wL41tS1cJ6HOEqiuPliD5owjGIw/RW3LxgAFZ1Pl6pHv25mYN967wPT2JFVEfpJlF4+gPqLN+UDGE3HIwnR2ZoH0EkWtVXDMtVYYhTdAhQs1p2O0QhuR8Wzj/Xp7LZZX6flBtDjJstVqlZa2SV0os+FczKD05uwqPU8rqk0vkqGON8W8j7YHG1dQlhhS7bLOJGOeNHXqdfGqT0Qmktuhh3+eXqXnZ0Gllz+wX+6vGuwI6bWOJEbf7oPsN/tR+H9uDUNN+RuQeQPadrLJBG2Gpr+X26bfKZdM5GR0Ldh1/iKR3OF8ZrN5hIh3wbQtL39qMPD8LRK5hKDhkQnim4iO/JzLSQEm4leyffbXoKLPxfjCchL/tlHtmqkH0/CbTORHXI8vNO2iTbZNz2D67alS6Gl08Ehx2rFA3N7M38lEwjCx+xlhb8n2gd2fRPV0xnoM/NsmXziy2fQT2LY7+37riyd2g91bM0mLtm3/BU1LmrE=",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},u=void 0,b={},L=[];function q(e){const s={p:"p",...(0,n.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"List Messages"}),"\n",(0,i.jsx)(o(),{method:"post",path:"/chat-svc/thread/{threadId}/messages",context:"endpoint"}),"\n",(0,i.jsx)(s.p,{children:"Fetch messages (and associated assets) for a specific chat thread."}),"\n",(0,i.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(c(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,i.jsx)(h(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(m(),{id:void 0,label:void 0,responses:{200:{description:"Messages and assets successfully retrieved",content:{"application/json":{schema:{properties:{messages:{items:{properties:{createdAt:{type:"string"},fileIds:{description:"FileIds defines the file attachments the message has.",items:{type:"string"},type:"array"},id:{example:"msg_emSOPlW58o",type:"string"},meta:{additionalProperties:!0,type:"object"},text:{description:'Text content of the message eg. "Hi, what\'s up?"',type:"string"},threadId:{description:"ThreadId of the message.",example:"thr_emSOeEUWAg",type:"string"},updatedAt:{type:"string"},userId:{description:"UserId is the id of the user who wrote the message.\nFor AI messages this field is empty.",type:"string"}},required:["id","threadId"],type:"object"},type:"array"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function G(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,i.jsx)(s,{...e,children:(0,i.jsx)(q,{...e})}):q(e)}}}]);