"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3097],{19579:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>f,contentTitle:()=>j,default:()=>g,frontMatter:()=>k,metadata:()=>n,toc:()=>w});const n=JSON.parse('{"id":"1backend/reset-password","title":"Reset Password","description":"Allows an administrator to change a user\'s password.","source":"@site/docs/1backend/reset-password.api.mdx","sourceDirName":"1backend","slug":"/1backend/reset-password","permalink":"/docs/1backend/reset-password","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"reset-password","title":"Reset Password","description":"Allows an administrator to change a user\'s password.","sidebar_label":"Reset Password","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VcGO2zYQ/RViLm0BWfJmG6DQqbtpULgNmsU6e9r4MKbGErMUyZCUHdfQvxcjWmt77QZpD7nYEjXkvDec92YHFQXplYvKGijhRmu7CQKNwKpVRoXoMVovohWyQVOTQNEF8j8E4TCEjfVVDhlYRx75iFkFJXgKFO/2nyEDhx5biuQDlI+7FxkfAnkx+w0yUPzqMDaQgcGWoARONeMjPH3ulKcKyug7yiDIhlqEcgdx6zgyRK9MDX2/SMEU4q2tthwhrYlkIj+ic1rJAWnxKXD+3dFRzjOPqCjwm6HNM4fzPBkE3dWXAGTjil1+Ihmh56VTzm9SKcfjxX0CfMaz54XgrAkJ0qvplP9OD3s+JV1QJUInJYWw6rTeQvbt9M9g9xn8fCnjzKxRq0r8MX//139JcFpf8t76byzggOTqHMmDwS421qu/qfpuSF5frkkkb1CLOfk1efF2OPP7QOJmJNl5FbeDwG4JPfmbLjZQPi5YERFr1l4S23wtYZFBS7GxLFdnh9YblFdCwZqbhLUsdkl9fTHoeeIOgg4Dx6TmzmveVWgrUTc2xPL1L6+ur4DTjqjmzDIRO8b2soYfto7Ex33IRxAry25ElVhuBYrgUJJAU4lon8gIlEkqYuVtK2JDYmQn3tlaGUGmclaZmI/e0hBW5A/ucrPvneFK4Lmy6NSftB1qzdd0f3CTt1+wdZrO3GG8ntEUjlxCmZUdTQjl0AfUouKaBdQUfg3K1J3G6K3JpW2P0N3NxLxzznq+nVTmJkZXFsXVEuUTmYo3FHBmLzezicGo1iRaJb3l21KSgnAa48r6lguilSQTEpV9vt/v3on1dT49yRbKothsNnltutz6utjvCwXWTk+u82nexFYzhki+De9X85TtADZssK7J58oWQ0jBlVaRywhXt4kIZMANleBP8+t8OvEy5y7KhvZs0Rwhved+FEcD5oT+keH/33G2b4RIX2LhNCrDOIaa7PYyeYRRJpBBeTSmTpSyyIAVweG73RIDPXjd97z8uSPPYl1ksEavcMnl4NmoAj9XUK5QB/oKtR/v96PiJ3EYoReBj21teBqsUXf8Bhk80fYwYvtFn40CYSDp45uUbsLKPNp85mR9Nu64kZJc/GrssffcvZ9/gAyW+1Hd2or3eNxwMXGTgNqB/mAfw9oONJq6w5pj05ksVtybykHGDCkbH5jVxVK8tIFEhH+Z1sUtu10yib5/jk+f/nXHs/ekaL6ZRd/3/wCSNUVt","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Checkout a git repository","permalink":"/docs/1backend/checkout-repo"},"next":{"title":"Change User Password","permalink":"/docs/1backend/change-password"}}');var r=t(74848),o=t(28453),a=t(53746),i=t.n(a),d=t(56518),c=t.n(d),l=t(99972),p=t.n(l),u=t(25342),h=t.n(u),m=(t(44215),t(82223),t(24861));const k={id:"reset-password",title:"Reset Password",description:"Allows an administrator to change a user's password.",sidebar_label:"Reset Password",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VcGO2zYQ/RViLm0BWfJmG6DQqbtpULgNmsU6e9r4MKbGErMUyZCUHdfQvxcjWmt77QZpD7nYEjXkvDec92YHFQXplYvKGijhRmu7CQKNwKpVRoXoMVovohWyQVOTQNEF8j8E4TCEjfVVDhlYRx75iFkFJXgKFO/2nyEDhx5biuQDlI+7FxkfAnkx+w0yUPzqMDaQgcGWoARONeMjPH3ulKcKyug7yiDIhlqEcgdx6zgyRK9MDX2/SMEU4q2tthwhrYlkIj+ic1rJAWnxKXD+3dFRzjOPqCjwm6HNM4fzPBkE3dWXAGTjil1+Ihmh56VTzm9SKcfjxX0CfMaz54XgrAkJ0qvplP9OD3s+JV1QJUInJYWw6rTeQvbt9M9g9xn8fCnjzKxRq0r8MX//139JcFpf8t76byzggOTqHMmDwS421qu/qfpuSF5frkkkb1CLOfk1efF2OPP7QOJmJNl5FbeDwG4JPfmbLjZQPi5YERFr1l4S23wtYZFBS7GxLFdnh9YblFdCwZqbhLUsdkl9fTHoeeIOgg4Dx6TmzmveVWgrUTc2xPL1L6+ur4DTjqjmzDIRO8b2soYfto7Ex33IRxAry25ElVhuBYrgUJJAU4lon8gIlEkqYuVtK2JDYmQn3tlaGUGmclaZmI/e0hBW5A/ucrPvneFK4Lmy6NSftB1qzdd0f3CTt1+wdZrO3GG8ntEUjlxCmZUdTQjl0AfUouKaBdQUfg3K1J3G6K3JpW2P0N3NxLxzznq+nVTmJkZXFsXVEuUTmYo3FHBmLzezicGo1iRaJb3l21KSgnAa48r6lguilSQTEpV9vt/v3on1dT49yRbKothsNnltutz6utjvCwXWTk+u82nexFYzhki+De9X85TtADZssK7J58oWQ0jBlVaRywhXt4kIZMANleBP8+t8OvEy5y7KhvZs0Rwhved+FEcD5oT+keH/33G2b4RIX2LhNCrDOIaa7PYyeYRRJpBBeTSmTpSyyIAVweG73RIDPXjd97z8uSPPYl1ksEavcMnl4NmoAj9XUK5QB/oKtR/v96PiJ3EYoReBj21teBqsUXf8Bhk80fYwYvtFn40CYSDp45uUbsLKPNp85mR9Nu64kZJc/GrssffcvZ9/gAyW+1Hd2or3eNxwMXGTgNqB/mAfw9oONJq6w5pj05ksVtybykHGDCkbH5jVxVK8tIFEhH+Z1sUtu10yib5/jk+f/nXHs/ekaL6ZRd/3/wCSNUVt",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},j=void 0,f={},w=[];function b(e){const s={p:"p",...(0,o.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Reset Password"}),"\n",(0,r.jsx)(i(),{method:"post",path:"/user-svc/{userId}/reset-password",context:"endpoint"}),"\n",(0,r.jsx)(s.p,{children:"Allows an administrator to change a user's password."}),"\n",(0,r.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,r.jsx)(c(),{parameters:[{description:"User ID",in:"path",name:"userId",required:!0,schema:{type:"string"}}]}),"\n",(0,r.jsx)(p(),{title:"Body",body:{content:{"application/json":{schema:{properties:{newPassword:{type:"string"},slug:{type:"string"}},type:"object"}}},description:"Change Password Request",required:!0}}),"\n",(0,r.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"Password changed successfully",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function g(e={}){const{wrapper:s}={...(0,o.R)(),...e.components};return s?(0,r.jsx)(s,{...e,children:(0,r.jsx)(b,{...e})}):b(e)}}}]);