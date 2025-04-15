"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5424],{76373:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>y,contentTitle:()=>m,default:()=>D,frontMatter:()=>b,metadata:()=>a,toc:()=>v});const a=JSON.parse('{"id":"1backend/get-thread","title":"Get Thread","description":"Fetch information about a specific chat thread by its ID","source":"@site/docs/1backend/get-thread.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-thread","permalink":"/docs/1backend/get-thread","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-thread","title":"Get Thread","description":"Fetch information about a specific chat thread by its ID","sidebar_label":"Get Thread","hide_title":true,"hide_table_of_contents":true,"api":"eJylVt9v2zYQ/lcIPm2AIjnLhgF6Wrp1gbti7eb0KTGGM3WW2FIkS57seob+9+FE2VEStVjRl8Qm78d3H+++81FWGFXQnrSzspS/I6lGaLt1oQU+E7BxHQkQ0aPSW62EaoAENQGhEpuD0BTF8jeZSecxDC7LSpayRrodbGQmPQRokTBEWd4dn2RMVimE5gMP1MhMWmhRljIlWnKYgB87HbCSJYUOMxlVgy3I8ijp4Nk2UtC2ln2/ZuPonY0Y+f6HxYL/zSaukECbKGKnFMa47Yw5iIAUNO6Q0ypnCS1xAPDeaDUUWbyPHOU4QeEDU0A65cRPOlKcoNs4ZxCs7LOxqOc+KiAQVtc0U1QmdZXiQuvNSM0/2K7+sv7VprmR2XMP0sSWzyrnY+G2ghocXzKfdXdeq2UVZyKMN6LCrbYYxb7RqhGDQ5yEFRs0ztZRkMvv7W26VmBFcF3dmIPYsCV/JsYDURDUUWxdGANEBqYJ2zhLyXgAIcCBv3e++gKBXcQwW867dDEA11U8UcP2XJobIEdkrDpOGPu/wPpp897xQ67PJm7zHhUNNk9P+OzHudZd2h0YXYlXqzd/fk2HPp2TlOByhhALHTUu6H+/bgTmEvw0XwFhsGDECsMOg3gZggvflqnPZETVBU2HQWZeIAQM1x01srxbsyZwbzH/v7KCrXaKX6FFatwoWINUsbksWOQu4k4V6bGL40mHeslpGHMSsy4Yti+MU2AaF6m8vLy6+llyvhOcFaNOEz4F9WyoDh7F/WhyL8XWGeP2OKgs6y8oFGArQe4DWgEqNZTYBtcO7cpdzGWJ167WVqCtvNOW8pOwNggVhgdpvR6feKD4Yf7B6z8wNS3vgUGXnCVQw6tgC5orjmAw/hK1rTsDFJzNlWsnsd8uxarz3gVmNZHUEPmyKC43oD6grdih4Fl5zML18sIC6R2KVqvgmGutMApvgHgrcTlGK7RxULZTvpu3r8XuKl88yhbLotjv93ltu9yFuhj9YgG1NxdX+SJvqDXDvGJo45vtKmV7ABv3UNcYcu2KwaSQZ1WVly9SITKT3A4J/iK/yhcXQeXcA5n0LlILdoL0BkmcN+Oj0o8P3f8te3h8RsJPVHgDetg4AyfHsb3v5Km95XkXZbI8r9p1JrmT2fB43EDEd8H0PR9/7DDwdK0zuYOgYcNE8ErXkT9XstyCifiFwr77exTC78V088+CPjWkPTDFYDr+JjP5AQ/TXwb9us9Ozc1g0vW1Uuhp4vhMT/rp+N+8vJWZhHEwH0aBg2WnDxx9FtTTUUoQ+G+ffcbleEyD1vdn+3T1WY/z/CZr5mjd9/1/NW1Xgw==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Delete a Thread","permalink":"/docs/1backend/delete-thread"},"next":{"title":"Update Thread","permalink":"/docs/1backend/update-thread"}}');var r=s(74848),d=s(28453),i=s(53746),n=s.n(i),c=s(56518),o=s.n(c),h=s(99972),p=s.n(h),l=s(25342),g=s.n(l),u=(s(44215),s(82223),s(24861));const b={id:"get-thread",title:"Get Thread",description:"Fetch information about a specific chat thread by its ID",sidebar_label:"Get Thread",hide_title:!0,hide_table_of_contents:!0,api:"eJylVt9v2zYQ/lcIPm2AIjnLhgF6Wrp1gbti7eb0KTGGM3WW2FIkS57seob+9+FE2VEStVjRl8Qm78d3H+++81FWGFXQnrSzspS/I6lGaLt1oQU+E7BxHQkQ0aPSW62EaoAENQGhEpuD0BTF8jeZSecxDC7LSpayRrodbGQmPQRokTBEWd4dn2RMVimE5gMP1MhMWmhRljIlWnKYgB87HbCSJYUOMxlVgy3I8ijp4Nk2UtC2ln2/ZuPonY0Y+f6HxYL/zSaukECbKGKnFMa47Yw5iIAUNO6Q0ypnCS1xAPDeaDUUWbyPHOU4QeEDU0A65cRPOlKcoNs4ZxCs7LOxqOc+KiAQVtc0U1QmdZXiQuvNSM0/2K7+sv7VprmR2XMP0sSWzyrnY+G2ghocXzKfdXdeq2UVZyKMN6LCrbYYxb7RqhGDQ5yEFRs0ztZRkMvv7W26VmBFcF3dmIPYsCV/JsYDURDUUWxdGANEBqYJ2zhLyXgAIcCBv3e++gKBXcQwW867dDEA11U8UcP2XJobIEdkrDpOGPu/wPpp897xQ67PJm7zHhUNNk9P+OzHudZd2h0YXYlXqzd/fk2HPp2TlOByhhALHTUu6H+/bgTmEvw0XwFhsGDECsMOg3gZggvflqnPZETVBU2HQWZeIAQM1x01srxbsyZwbzH/v7KCrXaKX6FFatwoWINUsbksWOQu4k4V6bGL40mHeslpGHMSsy4Yti+MU2AaF6m8vLy6+llyvhOcFaNOEz4F9WyoDh7F/WhyL8XWGeP2OKgs6y8oFGArQe4DWgEqNZTYBtcO7cpdzGWJ167WVqCtvNOW8pOwNggVhgdpvR6feKD4Yf7B6z8wNS3vgUGXnCVQw6tgC5orjmAw/hK1rTsDFJzNlWsnsd8uxarz3gVmNZHUEPmyKC43oD6grdih4Fl5zML18sIC6R2KVqvgmGutMApvgHgrcTlGK7RxULZTvpu3r8XuKl88yhbLotjv93ltu9yFuhj9YgG1NxdX+SJvqDXDvGJo45vtKmV7ABv3UNcYcu2KwaSQZ1WVly9SITKT3A4J/iK/yhcXQeXcA5n0LlILdoL0BkmcN+Oj0o8P3f8te3h8RsJPVHgDetg4AyfHsb3v5Km95XkXZbI8r9p1JrmT2fB43EDEd8H0PR9/7DDwdK0zuYOgYcNE8ErXkT9XstyCifiFwr77exTC78V088+CPjWkPTDFYDr+JjP5AQ/TXwb9us9Ozc1g0vW1Uuhp4vhMT/rp+N+8vJWZhHEwH0aBg2WnDxx9FtTTUUoQ+G+ffcbleEyD1vdn+3T1WY/z/CZr5mjd9/1/NW1Xgw==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,y={},v=[];function W(e){const t={p:"p",...(0,d.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"Get Thread"}),"\n",(0,r.jsx)(n(),{method:"get",path:"/chat-svc/thread/{threadId}",context:"endpoint"}),"\n",(0,r.jsx)(t.p,{children:"Fetch information about a specific chat thread by its ID"}),"\n",(0,r.jsx)(u.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,r.jsx)(o(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,r.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,r.jsx)(g(),{id:void 0,label:void 0,responses:{200:{description:"Thread details successfully retrieved",content:{"application/json":{schema:{properties:{exists:{type:"boolean"},thread:{properties:{createdAt:{type:"string"},id:{example:"thr_emSQnpJbhG",type:"string"},title:{description:"Title of the thread.",type:"string"},topicIds:{description:"TopicIds defines which topics the thread belongs to.\nTopics can roughly be thought of as tags for threads.",items:{type:"string"},type:"array"},updatedAt:{type:"string"},userIds:{description:"UserIds the ids of the users who can see this thread.",items:{type:"string"},type:"array"}},required:["id"],type:"object"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function D(e={}){const{wrapper:t}={...(0,d.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(W,{...e})}):W(e)}}}]);