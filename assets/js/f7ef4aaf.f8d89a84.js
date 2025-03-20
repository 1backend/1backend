"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9983],{97469:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>f,contentTitle:()=>j,default:()=>y,frontMatter:()=>m,metadata:()=>s,toc:()=>g});const s=JSON.parse('{"id":"1backend/list-container-logs","title":"List Logs","description":"List Container logs.","source":"@site/docs/1backend/list-container-logs.api.mdx","sourceDirName":"1backend","slug":"/1backend/list-container-logs","permalink":"/docs/1backend/list-container-logs","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"list-container-logs","title":"List Logs","description":"List Container logs.","sidebar_label":"List Logs","hide_title":true,"hide_table_of_contents":true,"api":"eJy9Vttu4zYQ/RVinhXJSbpAq6fmhtbdIA7izVNs7I6lscwNRTIkZa9r6N+LoXyN3QsWxb7YAufCc2bmjLSCknzhpA3SaMjhXvogbowOKDU5oUzl05Ee6Sd6a6QjL8KMxJdi43Dm50WuTJXPJS2+CEuult5Lo1NIwFhyyHn7JeSgpA/bxPem8pCAo7eGfLg25RLyFXBa0oEf0VolixidffUMbQW+mFGN/GQd5w6S/CYsZuV7VhCWliAHH5zUFbQJKFnLsGeROlBFjk3alHQyqk02J2bylYoALR+dqBUzEU8djzUj6aiEPLiGWj7w1mjfIb3o9fjvMM3gIyTfS537w/8yUO3/tTKH997sjEJ2jXW4EI0uyaml1JXYhov+bTrSd1Uqbk3xSm7PIstUPJiShDIFKm77Uf33uB3bHGGg8uq0VZ5u6K5rh4wijn450o+K0JPwRJFWaYqmJh1iTcXUuHhKeo5ODB7vHgZPN79/fhjc3n3u3x4TOB6F7QE6h8vTw9Im8NOpbvf1HJUsxR/DwcP3952cM+4/jm1Ecn6M5FljE2bGyT+p/GFIPpyuSSCnUYkhuTk5cRdz/hhIbQKeisbJsIT8ZQXXhI7cVRNmkL+M23ECAVljLzu9iOG8gHECNYWZ4dVmTdS+RQ6C7GA7ZqpbdT5S8/GOxin2y6JkZsaH/MPPF5fnwLdtwAyZXMdnH9L70n1aWhKjtcsIxNQoZRZUislSoPAWCxKoSxHMK2mBRbedxNSZOorg2Xd8eI9JLUiX1kgdWMeS888IS+JeaKy5clfrkYmd2GkFrfxInRK4O0+7vX73DWur6GgbbbqyXc+9na73N4Cemm1sEUeBapRcP4+K/K9e6qpRGJzRaWHqPaSPfTFsrDWOe9OVfBaCzbPsfILFK+mSAzI42utX/TONQc5J1LJwhjsnC/LCKgxT4+o0gi5I+0hrc99vj/difpn2Dm7zeZYtFou00k1qXJWt43yGlVVnl2kvnYVaxZVCrvaD6bC7bQfWL7CqyKXSZNEl46rLwCWF8+uOCCTAw9XB76WXae/MFenFL5yXh7NGvYd0+96Cd8z33sD/64fAekgCfQuZVSg144o1Wq1F8wIHmbjAjG+cAMuD7avVBD09O9W2fPzWkGPBjhOYo5M44Xq8jNtkM7Gss1dart9zpMMZS4ULharpRvbdRmmTTcRVUZAN/+i7L//HwfATJDBZf8XUpuQYhwv+HsAF5BA/hjg66jmerUChrhqs2LfLyerBtcp3umJIyeaBWW1MermH8L0uOyL8y7ROhqxWnWrbduvfmf42YrsMOm/u47ht278AVDCLTw==","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Check if Container Image is Pullable","permalink":"/docs/1backend/image-pullable"},"next":{"title":"Create a Generic Object","permalink":"/docs/1backend/create-object"}}');var i=n(74848),o=n(28453),a=n(53746),r=n.n(a),c=n(56518),d=n.n(c),l=n(99972),p=n.n(l),h=n(25342),b=n.n(h),u=(n(44215),n(82223),n(24861));const m={id:"list-container-logs",title:"List Logs",description:"List Container logs.",sidebar_label:"List Logs",hide_title:!0,hide_table_of_contents:!0,api:"eJy9Vttu4zYQ/RVinhXJSbpAq6fmhtbdIA7izVNs7I6lscwNRTIkZa9r6N+LoXyN3QsWxb7YAufCc2bmjLSCknzhpA3SaMjhXvogbowOKDU5oUzl05Ee6Sd6a6QjL8KMxJdi43Dm50WuTJXPJS2+CEuult5Lo1NIwFhyyHn7JeSgpA/bxPem8pCAo7eGfLg25RLyFXBa0oEf0VolixidffUMbQW+mFGN/GQd5w6S/CYsZuV7VhCWliAHH5zUFbQJKFnLsGeROlBFjk3alHQyqk02J2bylYoALR+dqBUzEU8djzUj6aiEPLiGWj7w1mjfIb3o9fjvMM3gIyTfS537w/8yUO3/tTKH997sjEJ2jXW4EI0uyaml1JXYhov+bTrSd1Uqbk3xSm7PIstUPJiShDIFKm77Uf33uB3bHGGg8uq0VZ5u6K5rh4wijn450o+K0JPwRJFWaYqmJh1iTcXUuHhKeo5ODB7vHgZPN79/fhjc3n3u3x4TOB6F7QE6h8vTw9Im8NOpbvf1HJUsxR/DwcP3952cM+4/jm1Ecn6M5FljE2bGyT+p/GFIPpyuSSCnUYkhuTk5cRdz/hhIbQKeisbJsIT8ZQXXhI7cVRNmkL+M23ECAVljLzu9iOG8gHECNYWZ4dVmTdS+RQ6C7GA7ZqpbdT5S8/GOxin2y6JkZsaH/MPPF5fnwLdtwAyZXMdnH9L70n1aWhKjtcsIxNQoZRZUislSoPAWCxKoSxHMK2mBRbedxNSZOorg2Xd8eI9JLUiX1kgdWMeS888IS+JeaKy5clfrkYmd2GkFrfxInRK4O0+7vX73DWur6GgbbbqyXc+9na73N4Cemm1sEUeBapRcP4+K/K9e6qpRGJzRaWHqPaSPfTFsrDWOe9OVfBaCzbPsfILFK+mSAzI42utX/TONQc5J1LJwhjsnC/LCKgxT4+o0gi5I+0hrc99vj/difpn2Dm7zeZYtFou00k1qXJWt43yGlVVnl2kvnYVaxZVCrvaD6bC7bQfWL7CqyKXSZNEl46rLwCWF8+uOCCTAw9XB76WXae/MFenFL5yXh7NGvYd0+96Cd8z33sD/64fAekgCfQuZVSg144o1Wq1F8wIHmbjAjG+cAMuD7avVBD09O9W2fPzWkGPBjhOYo5M44Xq8jNtkM7Gss1dart9zpMMZS4ULharpRvbdRmmTTcRVUZAN/+i7L//HwfATJDBZf8XUpuQYhwv+HsAF5BA/hjg66jmerUChrhqs2LfLyerBtcp3umJIyeaBWW1MermH8L0uOyL8y7ROhqxWnWrbduvfmf42YrsMOm/u47ht278AVDCLTw==",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},j=void 0,f={},g=[];function L(e){const t={code:"code",p:"p",...(0,o.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"List Logs"}),"\n",(0,i.jsx)(r(),{method:"post",path:"/container-svc/logs",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"List Container logs."}),"\n",(0,i.jsxs)(t.p,{children:["Requires the ",(0,i.jsx)(t.code,{children:"container-svc:log:view"})," permission."]}),"\n",(0,i.jsx)(u.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(d(),{parameters:void 0}),"\n",(0,i.jsx)(p(),{title:"Body",body:{content:{"application/json":{schema:{properties:{containerId:{type:"string"},limit:{type:"integer"},nodeId:{type:"string"}},type:"object"}}},description:"List Logs Request",required:!0}}),"\n",(0,i.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{logs:{items:{properties:{containerId:{description:"ContainerId is the raw underlying container ID.\nEg. Docker container id. Node local.",type:"string"},content:{type:"string"},createdAt:{type:"string"},id:{type:"string"},nodeId:{description:"Node Id\nPlease see the documentation for the envar OPENORCH_NODE_ID",type:"string"}},type:"object"},type:"array"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function y(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(L,{...e})}):L(e)}}}]);