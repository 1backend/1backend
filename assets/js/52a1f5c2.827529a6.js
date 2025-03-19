"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7908],{99914:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>T,contentTitle:()=>b,default:()=>f,frontMatter:()=>m,metadata:()=>a,toc:()=>v});const a=JSON.parse('{"id":"1backend/get-threads","title":"Get Threads","description":"Fetch all chat threads associated with a specific user","source":"@site/docs/1backend/get-threads.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-threads","permalink":"/docs/1backend/get-threads","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-threads","title":"Get Threads","description":"Fetch all chat threads associated with a specific user","sidebar_label":"Get Threads","hide_title":true,"hide_table_of_contents":true,"api":"eJylVk1z2zYQ/SuYPdOkP5qZlqfKmdTjNFOnkXOyNZ0VuCIRgwACgFJUDf97ZwnSViQ1k7QXmwJ2sW8f3j5yBxUF6ZWLyhoo4TeKshGotZANRhEbT1gFgSFYqTBSJTYqNgJFcCTVSknRBfKQgXXkkQ+5raCEmuJ9SoUMPH3uKMRrW22h3IG0JpKJ/IjOaSWHtOJTYAA7CLKhFvkpbh1BCXb5iWSEvu+zA7A3FMVYRnxIRaDnesFZEyjwIZfn5/zv68QpKXRSUgirTuut8BS9ojVVkH0/Rue58ahSsZEuflSR2nAcIT0xi7O412CIXpmakauKl+kLtk7zTmz8X9TO/zTu7bK5gew4I6rIkUcN8rKwKxEbGi8xP5lunZK3CfHBCeOOqGilDAWxaZRsxJAQ9o4VS9LW1EFEmz+a+7Qt0Qhvu7rRW7HkSH6OjAeDiFgHsbJ+EhcDe2brGGFaQO9xy787V32DQBbjyXY+po0BuKrCRA3Hc2t2gByIsaqwx9j3AuuTzJWnCsoHvshFdijgEzlHEu8z+OmUYm/NGrWqxNv53R8/os8D2GOBixMEGexiY736+8cG4FSBV6c7iOQNajEnvyYv3nhv/f+r1GcQSHZexS2UDzu4JvTkZ11soHxY9HwDWAe+j9dsZvO15FtpKTaWTcrZECEDhxwPBRveWVjLIj47VxighuHwzmuOKrSVqBsbYvnq58urC+AyE4o5g02Dvo/laLa2jsTjGPIIYmW1thuqxHI7OCtKEmgqEe0TGYEy6UqsvG0H1bKYuRvxztbKCDKVs8rEQa98fkNYDaZssGXKZuPNDsy+2AA69TslHTLbH15s+s3kQDs2JbOyk2+jHC6KWlTMRkBN4degTN1pjN6aXNp2r+77WzHvnLOeeU4ENjG6siguliifyFScUMCRs89EZ9RKMSUpTlS0Jm1dSyYKpzGurG8HF2mV9JYvSkkKZ0sMVInZrdjT0uAwWkkyYbDKCd3N+3difZWff4UtlEWx2Wzy2nS59XUx5oUCa6fPrvLzvImtHoaZfBvuVvNU+qW1sMG6Jp8rWwwhBTzbNFxcp3YgAxZWavY8v8rPz7zML3/hc1mWLZo9pHvvOThgau91+t/f3aMeIn2JhdOozGClTMlunI4HmKaDw0coiwx4Dnh3t2PiP3rd97z8uSPPI7nIYI1e4ZKbf1j02SRNHqgn2kIJrxP8M54JZgV1l7R5YAV9NmXMpCQXvxm7P+Xv7+b3kMFy/PpobcU5Hjf8ZYIbKGH4ehmEwgHD2g40mrrDmmPTmTwmOI7zywAxpGx64K6mLbPdQ3g4gKkR/sttnUzZ7dJ49v1zfNr614znqU/RfIuLvu//Ad8TaCI=","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"List Messages","permalink":"/docs/1backend/get-messages"},"next":{"title":"Get Config","permalink":"/docs/1backend/get-config"}}');var i=s(74848),n=s(28453),r=s(53746),d=s.n(r),o=s(56518),c=s.n(o),p=s(99972),h=s.n(p),u=s(25342),l=s.n(u),g=(s(44215),s(82223),s(24861));const m={id:"get-threads",title:"Get Threads",description:"Fetch all chat threads associated with a specific user",sidebar_label:"Get Threads",hide_title:!0,hide_table_of_contents:!0,api:"eJylVk1z2zYQ/SuYPdOkP5qZlqfKmdTjNFOnkXOyNZ0VuCIRgwACgFJUDf97ZwnSViQ1k7QXmwJ2sW8f3j5yBxUF6ZWLyhoo4TeKshGotZANRhEbT1gFgSFYqTBSJTYqNgJFcCTVSknRBfKQgXXkkQ+5raCEmuJ9SoUMPH3uKMRrW22h3IG0JpKJ/IjOaSWHtOJTYAA7CLKhFvkpbh1BCXb5iWSEvu+zA7A3FMVYRnxIRaDnesFZEyjwIZfn5/zv68QpKXRSUgirTuut8BS9ojVVkH0/Rue58ahSsZEuflSR2nAcIT0xi7O412CIXpmakauKl+kLtk7zTmz8X9TO/zTu7bK5gew4I6rIkUcN8rKwKxEbGi8xP5lunZK3CfHBCeOOqGilDAWxaZRsxJAQ9o4VS9LW1EFEmz+a+7Qt0Qhvu7rRW7HkSH6OjAeDiFgHsbJ+EhcDe2brGGFaQO9xy787V32DQBbjyXY+po0BuKrCRA3Hc2t2gByIsaqwx9j3AuuTzJWnCsoHvshFdijgEzlHEu8z+OmUYm/NGrWqxNv53R8/os8D2GOBixMEGexiY736+8cG4FSBV6c7iOQNajEnvyYv3nhv/f+r1GcQSHZexS2UDzu4JvTkZ11soHxY9HwDWAe+j9dsZvO15FtpKTaWTcrZECEDhxwPBRveWVjLIj47VxighuHwzmuOKrSVqBsbYvnq58urC+AyE4o5g02Dvo/laLa2jsTjGPIIYmW1thuqxHI7OCtKEmgqEe0TGYEy6UqsvG0H1bKYuRvxztbKCDKVs8rEQa98fkNYDaZssGXKZuPNDsy+2AA69TslHTLbH15s+s3kQDs2JbOyk2+jHC6KWlTMRkBN4degTN1pjN6aXNp2r+77WzHvnLOeeU4ENjG6siguliifyFScUMCRs89EZ9RKMSUpTlS0Jm1dSyYKpzGurG8HF2mV9JYvSkkKZ0sMVInZrdjT0uAwWkkyYbDKCd3N+3difZWff4UtlEWx2Wzy2nS59XUx5oUCa6fPrvLzvImtHoaZfBvuVvNU+qW1sMG6Jp8rWwwhBTzbNFxcp3YgAxZWavY8v8rPz7zML3/hc1mWLZo9pHvvOThgau91+t/f3aMeIn2JhdOozGClTMlunI4HmKaDw0coiwx4Dnh3t2PiP3rd97z8uSPPI7nIYI1e4ZKbf1j02SRNHqgn2kIJrxP8M54JZgV1l7R5YAV9NmXMpCQXvxm7P+Xv7+b3kMFy/PpobcU5Hjf8ZYIbKGH4ehmEwgHD2g40mrrDmmPTmTwmOI7zywAxpGx64K6mLbPdQ3g4gKkR/sttnUzZ7dJ49v1zfNr614znqU/RfIuLvu//Ad8TaCI=",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},b=void 0,T={},v=[];function k(e){const t={p:"p",...(0,n.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Get Threads"}),"\n",(0,i.jsx)(d(),{method:"post",path:"/chat-svc/threads",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Fetch all chat threads associated with a specific user"}),"\n",(0,i.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(c(),{parameters:void 0}),"\n",(0,i.jsx)(h(),{title:"Body",body:{content:{"application/json":{schema:{type:"object"}}},description:"Get Threads Request"}}),"\n",(0,i.jsx)(l(),{id:void 0,label:void 0,responses:{200:{description:"Threads successfully retrieved",content:{"application/json":{schema:{properties:{threads:{items:{properties:{createdAt:{type:"string"},id:{example:"thr_emSQnpJbhG",type:"string"},title:{description:"Title of the thread.",type:"string"},topicIds:{description:"TopicIds defines which topics the thread belongs to.\nTopics can roughly be thought of as tags for threads.",items:{type:"string"},type:"array"},updatedAt:{type:"string"},userIds:{description:"UserIds the ids of the users who can see this thread.",items:{type:"string"},type:"array"}},required:["id"],type:"object"},type:"array"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function f(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(k,{...e})}):k(e)}}}]);