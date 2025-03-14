"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6685],{64323:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>v,contentTitle:()=>m,default:()=>I,frontMatter:()=>k,metadata:()=>s,toc:()=>y});const s=JSON.parse('{"id":"1backend/login","title":"Login","description":"Authenticates a user and returns a token.","source":"@site/docs/1backend/login.api.mdx","sourceDirName":"1backend","slug":"/1backend/login","permalink":"/docs/1backend/login","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"login","title":"Login","description":"Authenticates a user and returns a token.","sidebar_label":"Login","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VsFuGzcQ/RWCl15Wu7LdAO2eKhdB4daojShBD5YPI+5olzGXZIakFFUQkI/IF+ZLguGuYjlWgzaH6qLFcGY48+bN293JBoMi7aN2VtZylmKHNmoFEYMAkQKSANsIwpjIsim6B7SlLKTzSMBxV42spXGttrKQhO8Shnjpmq2sd1I5G9FGfgTvDSfWzlZvA1+3k0F12AM/eeJ8UWM4hIHKYXHrUdYyRNK2lftCeghh46g5eRhMak8c7IuDxS3foopyz6anvV9zB+LVUP/YiSZsZB0p4Z4NwTsbhgrPp1P+O5UiJKUwhFUysvheADLKz82gol7j84tn2T7MJoiMnrYidih6F6JIfhLdpIGIQtuVoz4XUS7sXx3accw/BEHOoFAd2BbDpw8fm8QZByuEoFvbo42FcNSC1X/nHAurCIeWjl0wqvLTh49gjMD3OkRt20NxQCh6oAdshLZDP+XCPm2AfQhTwEY0iScoMr0Ksem0wS9xo/vCDv7cMzZi5Uh0OkRHWoERhCsktAqZsyMJls4ZBMuEyeVjMzvNtQYN/vOpPs3BL7N7dpJ8843LeApXp1KeYO8JPu8L+eMpUl7ZNRjdiN/nN39+PyGRyNG/3Kx9IV+criQiWTBijrRGEi9zzv+nJDZCG2R9J9+wqM3XSt4XssfYOdYv7/LSe4idrGXFs5iEtaoOuhZyyRy/k4kM+1TGKTCdC7F+8dP5xZnc37OfSqTjds5FD3VeIhASS+tzSF5vPYrF6LKQYuWMcRtsxHIrQAQPCrP+Zk4JUIMkiRW5Pq/3oRcxSA/axjttI3Ndc/4OoUHG2EKPo8A7Gpf3cSHA6z9wm6Fj1F89ivjL99B7g09E+YD0sRY/2gYJPl4Tu3JfiTr2oBnDAAbDL0HbNhmI5GypXH9U7e2VmCfvHfFsBti7GH1dVWdLUA9oGw6o5DMxnwnl+t5ZMbplVdi6RGJ2JY5IxjrXa0WO56sVhskSWHYY82XSJrICBgUmy4fRCm3IYBwq/O32WqwvyumT+kJdVZvNpmxtKh211RgXKmi9mVyU07KLvclagdSHm9V8uP2xvbCBtkUqtauyS8Wz0pEHIc8uh55kIZmSQ8PT8qKcTkiV5z/nl6QLsQd7VOn1SOQnOB29oP/Tu3+kTcT3sfIGdJbS3P9uXKE7eVghxi3ffV9IXhY+2+0Y5jdk9ns2v0tIW1nf3RdyDaRhyX3e3e+LA3956x5wK2v561DvhBeHAQCTBgJ/pRv74hAxUwp9/KbvsRDc3sxfy0Iuxw+Y3jUcQ7DhTwLYyFrmb59MHnbItp00YNsELfsOOfn3GaE2Rk0=","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Save Grants","permalink":"/docs/1backend/save-grants"},"next":{"title":"Create an Organization","permalink":"/docs/1backend/create-organization"}}');var i=n(74848),o=n(28453),a=n(53746),r=n.n(a),d=n(56518),c=n.n(d),p=n(99972),l=n.n(p),h=n(25342),u=n.n(h),g=(n(44215),n(82223),n(24861));const k={id:"login",title:"Login",description:"Authenticates a user and returns a token.",sidebar_label:"Login",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VsFuGzcQ/RWCl15Wu7LdAO2eKhdB4daojShBD5YPI+5olzGXZIakFFUQkI/IF+ZLguGuYjlWgzaH6qLFcGY48+bN293JBoMi7aN2VtZylmKHNmoFEYMAkQKSANsIwpjIsim6B7SlLKTzSMBxV42spXGttrKQhO8Shnjpmq2sd1I5G9FGfgTvDSfWzlZvA1+3k0F12AM/eeJ8UWM4hIHKYXHrUdYyRNK2lftCeghh46g5eRhMak8c7IuDxS3foopyz6anvV9zB+LVUP/YiSZsZB0p4Z4NwTsbhgrPp1P+O5UiJKUwhFUysvheADLKz82gol7j84tn2T7MJoiMnrYidih6F6JIfhLdpIGIQtuVoz4XUS7sXx3accw/BEHOoFAd2BbDpw8fm8QZByuEoFvbo42FcNSC1X/nHAurCIeWjl0wqvLTh49gjMD3OkRt20NxQCh6oAdshLZDP+XCPm2AfQhTwEY0iScoMr0Ksem0wS9xo/vCDv7cMzZi5Uh0OkRHWoERhCsktAqZsyMJls4ZBMuEyeVjMzvNtQYN/vOpPs3BL7N7dpJ8843LeApXp1KeYO8JPu8L+eMpUl7ZNRjdiN/nN39+PyGRyNG/3Kx9IV+criQiWTBijrRGEi9zzv+nJDZCG2R9J9+wqM3XSt4XssfYOdYv7/LSe4idrGXFs5iEtaoOuhZyyRy/k4kM+1TGKTCdC7F+8dP5xZnc37OfSqTjds5FD3VeIhASS+tzSF5vPYrF6LKQYuWMcRtsxHIrQAQPCrP+Zk4JUIMkiRW5Pq/3oRcxSA/axjttI3Ndc/4OoUHG2EKPo8A7Gpf3cSHA6z9wm6Fj1F89ivjL99B7g09E+YD0sRY/2gYJPl4Tu3JfiTr2oBnDAAbDL0HbNhmI5GypXH9U7e2VmCfvHfFsBti7GH1dVWdLUA9oGw6o5DMxnwnl+t5ZMbplVdi6RGJ2JY5IxjrXa0WO56sVhskSWHYY82XSJrICBgUmy4fRCm3IYBwq/O32WqwvyumT+kJdVZvNpmxtKh211RgXKmi9mVyU07KLvclagdSHm9V8uP2xvbCBtkUqtauyS8Wz0pEHIc8uh55kIZmSQ8PT8qKcTkiV5z/nl6QLsQd7VOn1SOQnOB29oP/Tu3+kTcT3sfIGdJbS3P9uXKE7eVghxi3ffV9IXhY+2+0Y5jdk9ns2v0tIW1nf3RdyDaRhyX3e3e+LA3956x5wK2v561DvhBeHAQCTBgJ/pRv74hAxUwp9/KbvsRDc3sxfy0Iuxw+Y3jUcQ7DhTwLYyFrmb59MHnbItp00YNsELfsOOfn3GaE2Rk0=",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,v={},y=[];function f(e){const t={p:"p",...(0,o.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Login"}),"\n",(0,i.jsx)(r(),{method:"post",path:"/user-svc/login",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Authenticates a user and returns a token."}),"\n",(0,i.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(c(),{parameters:void 0}),"\n",(0,i.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{contact:{type:"string"},password:{type:"string"},slug:{type:"string"}},type:"object"}}},description:"Login Request",required:!0}}),"\n",(0,i.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"Login successful",content:{"application/json":{schema:{properties:{token:{properties:{active:{description:"Active tokens contain the most up-to-date information.\nWhen a user's role changes\u2014due to role assignment, organization\ncreation/assignment, etc.\u2014all existing tokens are marked inactive.\nActive tokens are reused during login, while inactive tokens\nare retained for historical reference.",type:"boolean"},createdAt:{type:"string"},deletedAt:{type:"string"},id:{type:"string"},token:{type:"string"},updatedAt:{type:"string"},userId:{type:"string"}},type:"object"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function I(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(f,{...e})}):f(e)}}}]);