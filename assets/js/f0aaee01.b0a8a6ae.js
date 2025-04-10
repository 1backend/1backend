"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[516],{76268:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>x,contentTitle:()=>X,default:()=>k,frontMatter:()=>m,metadata:()=>r,toc:()=>K});const r=JSON.parse('{"id":"1backend/subscribe-to-prompt-responses","title":"Subscribe to Prompt Responses by Thread","description":"Subscribe to prompt responses by thread via Server-Sent Events (SSE).","source":"@site/docs/1backend/subscribe-to-prompt-responses.api.mdx","sourceDirName":"1backend","slug":"/1backend/subscribe-to-prompt-responses","permalink":"/docs/1backend/subscribe-to-prompt-responses","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"subscribe-to-prompt-responses","title":"Subscribe to Prompt Responses by Thread","description":"Subscribe to prompt responses by thread via Server-Sent Events (SSE).","sidebar_label":"Subscribe to Prompt Responses by Thread","hide_title":true,"hide_table_of_contents":true,"api":"eJzdVU2P2zYQ/SvEnJKAK9l1iwI6ddMuAqMpulh5D8WugdDUWGJXIhmSsusI+u/FiJbsdXaTUy+92JQ0H2/ezBt2UKCXTtmgjIYM8nZDjxtkwTDrTGMDc+it0R492xxYqByKgu2UYDm6HbqrHHVgNzvUwbM3eX7zNnnUf5mWSaGZPw8XXT3b4NY4ZKHCAxMOmXQoAhbJo15VyHxwKBos6KB06QcTs2XhYJF9yoevv1atfvrEmcchDLsdgK4OFj1DXVijdGBb41hDiQoMQtU+AQ7GohNU6rKADCZ0KxMj3I2VAgcrnGgwoPOQPXQXNK0iC8vfgIOiF1aECjho0SBkECtdFsDB4edWOSwgC65FDl5W2AjIOqCCCMRQJvT9mozH/FkHP8xm9HfRn6F+pcupK8BBGh1QB7J+l76jv9ez9D2HH1+K/IfynuKO2NlEwHcyWEesBhVRo3PGvZSYj2/M5m+UYYIy/xrKvRZtqIxTX7D475L3HDzK1qlwGDr8HoVDd92GCrKHNbUjiJKaD3E4WL6TsObQYKgMjU+JYZgTcoA0iuXK7+Tx6NNu5LJPp86m09AB5ScFxQFrXU1h0tpIUVfGh2w+Xyx+BgIy4syp7ljqOdqvppOk8ng0eQS2NXVt9liQfgXzVkhkQhcsmCfUTMg4oWzrTDPI6d6jo2rZR1MqPSkqGYe9QlEMY3Ec9+tjtwZlwcS1sOp3PAzsK701hJNaKeTQSmyEooq9qNH/QqPX1iI4oxNpmrPYt0uWt9YaR2RHkqoQbJam842QT6gLckih5xcsXC+vtAhqh6xR0hniWkn0zNYibI1rqJxaSSQNZd2U78PtR7ZbJLNn2XyWpvv9Pil1mxhXpkc/n4rS1leLZJZUoakJQ0DX+D+3ecx2Auv3oizRJcqkg0lKPKlQk8n8fSwEONA4RPizZJHMrpxMFj9RXGt8aIQ+Q/psVR9H9O58VcclBRe8dCc9/d+3/XEOA/4TUlsLpYnIoandUbYPcJItaTkKFzhkz1b46VY4qXfNgVRKMbpuIzzeu7rv6fXnFh2tlDWHnXBKbKjJdIUoT+cCsq2oPX6jL2/ujrfGW3Z+07xYzyg2faDxEXVLT8DhCQ/nN1G/7vkoXAITP19LiTacOdKC7c+X3IebFXAQxz1zUjb58/FAAV/EcbkZYlb67fkrLl0X90bfT/bx06se0zqK1kTLuu/7fwGZ5w0U","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"List Prompts","permalink":"/docs/1backend/list-prompts"},"next":{"title":"Remove Prompt","permalink":"/docs/1backend/remove-prompt"}}');var o=t(74848),n=t(28453),i=t(53746),p=t.n(i),a=t(56518),c=t.n(a),d=t(99972),b=t.n(d),h=t(25342),l=t.n(h),u=(t(44215),t(82223),t(24861));const m={id:"subscribe-to-prompt-responses",title:"Subscribe to Prompt Responses by Thread",description:"Subscribe to prompt responses by thread via Server-Sent Events (SSE).",sidebar_label:"Subscribe to Prompt Responses by Thread",hide_title:!0,hide_table_of_contents:!0,api:"eJzdVU2P2zYQ/SvEnJKAK9l1iwI6ddMuAqMpulh5D8WugdDUWGJXIhmSsusI+u/FiJbsdXaTUy+92JQ0H2/ezBt2UKCXTtmgjIYM8nZDjxtkwTDrTGMDc+it0R492xxYqByKgu2UYDm6HbqrHHVgNzvUwbM3eX7zNnnUf5mWSaGZPw8XXT3b4NY4ZKHCAxMOmXQoAhbJo15VyHxwKBos6KB06QcTs2XhYJF9yoevv1atfvrEmcchDLsdgK4OFj1DXVijdGBb41hDiQoMQtU+AQ7GohNU6rKADCZ0KxMj3I2VAgcrnGgwoPOQPXQXNK0iC8vfgIOiF1aECjho0SBkECtdFsDB4edWOSwgC65FDl5W2AjIOqCCCMRQJvT9mozH/FkHP8xm9HfRn6F+pcupK8BBGh1QB7J+l76jv9ez9D2HH1+K/IfynuKO2NlEwHcyWEesBhVRo3PGvZSYj2/M5m+UYYIy/xrKvRZtqIxTX7D475L3HDzK1qlwGDr8HoVDd92GCrKHNbUjiJKaD3E4WL6TsObQYKgMjU+JYZgTcoA0iuXK7+Tx6NNu5LJPp86m09AB5ScFxQFrXU1h0tpIUVfGh2w+Xyx+BgIy4syp7ljqOdqvppOk8ng0eQS2NXVt9liQfgXzVkhkQhcsmCfUTMg4oWzrTDPI6d6jo2rZR1MqPSkqGYe9QlEMY3Ec9+tjtwZlwcS1sOp3PAzsK701hJNaKeTQSmyEooq9qNH/QqPX1iI4oxNpmrPYt0uWt9YaR2RHkqoQbJam842QT6gLckih5xcsXC+vtAhqh6xR0hniWkn0zNYibI1rqJxaSSQNZd2U78PtR7ZbJLNn2XyWpvv9Pil1mxhXpkc/n4rS1leLZJZUoakJQ0DX+D+3ecx2Auv3oizRJcqkg0lKPKlQk8n8fSwEONA4RPizZJHMrpxMFj9RXGt8aIQ+Q/psVR9H9O58VcclBRe8dCc9/d+3/XEOA/4TUlsLpYnIoandUbYPcJItaTkKFzhkz1b46VY4qXfNgVRKMbpuIzzeu7rv6fXnFh2tlDWHnXBKbKjJdIUoT+cCsq2oPX6jL2/ujrfGW3Z+07xYzyg2faDxEXVLT8DhCQ/nN1G/7vkoXAITP19LiTacOdKC7c+X3IebFXAQxz1zUjb58/FAAV/EcbkZYlb67fkrLl0X90bfT/bx06se0zqK1kTLuu/7fwGZ5w0U",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},X=void 0,x={},K=[];function S(e){const s={code:"code",p:"p",...(0,n.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"Subscribe to Prompt Responses by Thread"}),"\n",(0,o.jsx)(p(),{method:"get",path:"/prompt-svc/prompts/{threadId}/responses/subscribe",context:"endpoint"}),"\n",(0,o.jsxs)(s.p,{children:["Subscribe to prompt responses by thread via Server-Sent Events (SSE).\nYou can subscribe to threads before they are created.\nThe streamed strings are of type ",(0,o.jsx)(s.code,{children:"StreamChunk"}),", see the PromptTypes endpoint for more details."]}),"\n",(0,o.jsx)(u.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,o.jsx)(c(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,o.jsx)(b(),{title:"Body",body:void 0}),"\n",(0,o.jsx)(l(),{id:void 0,label:void 0,responses:{200:{description:"Streaming response",content:{"*/*":{schema:{type:"string"}}}},400:{description:"Missing threadId parameter",content:{"*/*":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"*/*":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function k(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,o.jsx)(s,{...e,children:(0,o.jsx)(S,{...e})}):S(e)}}}]);