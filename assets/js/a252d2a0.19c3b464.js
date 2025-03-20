"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6830],{70875:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>m,contentTitle:()=>B,default:()=>x,frontMatter:()=>k,metadata:()=>o,toc:()=>D});const o=JSON.parse('{"id":"1backend/delete-role","title":"Delete a Role","description":"Delete a role based on the role ID.","source":"@site/docs/1backend/delete-role.api.mdx","sourceDirName":"1backend","slug":"/1backend/delete-role","permalink":"/docs/1backend/delete-role","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"delete-role","title":"Delete a Role","description":"Delete a role based on the role ID.","sidebar_label":"Delete a Role","hide_title":true,"hide_table_of_contents":true,"api":"eJylVE1v2zgQ/SvEnHYBRnKSLbDVqSkSFG6DTRE3p9SHMTW22FAkS1L2egX+92IsOXY+ulhsL7Ykzsx78/hmeqgpqqB90s5CBZdkKJFAEZwhscBItXBWpIaGL9PLAiQ4TwE5Y1pDBfUu59YZAgkeA7aUKESo7vtn1W+HEiBB86vH1IAEiy1BBVx/WoOEQN87HaiGKoWOJETVUItQ9ZC2niNjCtquIOc5B0fvbKTI52eTCf89Bb35BBKUs4ls4lP03mi1o19+ixzSv4Rwi2+kEuScs4Q/Xis7tWs0uhYfZzd//Q+Axx4GgNOXAHcWu9S4oP+h+lcB3rzeQaJg0YgZhTUFcRWCC7+GlCVEUl3Qabu7/veEgcJFlxqo7ud8XwlX7Ay4ixTEbK1gLqGl1LiDk3Yu4gwou0jhJK5VyeYo+8EiGRiFKQ8e64Lh2NI4haZxMVVv/jw7PwWG27OZMenBJMecnkvyZetJfB1DvoJYOmPchmqx2AoU0aMigbYWyT2QFagGo4plcO1uRvZdiWu30laQrb3TNhV7xzeENYWD5y/GG94pDHKvKHr9ibbAcmq7dMyTLwXV7lKoRc0dRzQU30VtV53BFJwtlGuPan+eilnnvQsJ5ChSk5KvyvJ0geqBbM0JJWT5TIWL6YnFpNckWq2CY621oii8wbR0oeV2jFZkIzGfPd6Hz9difV5MnqDFqiw3m02xsl3hwqoc82KJK29OzotJ0aTWMIdEoY03y9mAdiAbN7haUSi0K3chJeukk+GQ0/dDIyCB7TDQnxTnxeQkqOLsLdf1LqYW7RHTxyU37qwn3fcH///HbTheWqK/U+kNasuoOwX60cj3sDcyr7cBtBr33VwCe5aD+p4h7oLJmT9/7yjwGM0lrDFoXHDLvFN15OcaqiWaSP/C/7fbcZX+Lg6r91W6e+PZLUuJpuM3kPBA28NqzvMs9xZmIsPhhVLk01Hai6WRj2f88ur66ssVSMBxAg+e53py/8AAr7J6PjMDC/7N8icpfT9MVM6P8cPRTzMeB3WIZpHmOecfJcJtbA==","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Create a New Role","permalink":"/docs/1backend/create-role"},"next":{"title":"Get Permissions by Role","permalink":"/docs/1backend/get-permissions-by-role"}}');var n=i(74848),a=i(28453),s=i(53746),l=i.n(s),c=i(56518),d=i.n(c),r=i(99972),p=i.n(r),h=i(25342),u=i.n(h),b=(i(44215),i(82223),i(24861));const k={id:"delete-role",title:"Delete a Role",description:"Delete a role based on the role ID.",sidebar_label:"Delete a Role",hide_title:!0,hide_table_of_contents:!0,api:"eJylVE1v2zgQ/SvEnHYBRnKSLbDVqSkSFG6DTRE3p9SHMTW22FAkS1L2egX+92IsOXY+ulhsL7Ykzsx78/hmeqgpqqB90s5CBZdkKJFAEZwhscBItXBWpIaGL9PLAiQ4TwE5Y1pDBfUu59YZAgkeA7aUKESo7vtn1W+HEiBB86vH1IAEiy1BBVx/WoOEQN87HaiGKoWOJETVUItQ9ZC2niNjCtquIOc5B0fvbKTI52eTCf89Bb35BBKUs4ls4lP03mi1o19+ixzSv4Rwi2+kEuScs4Q/Xis7tWs0uhYfZzd//Q+Axx4GgNOXAHcWu9S4oP+h+lcB3rzeQaJg0YgZhTUFcRWCC7+GlCVEUl3Qabu7/veEgcJFlxqo7ud8XwlX7Ay4ixTEbK1gLqGl1LiDk3Yu4gwou0jhJK5VyeYo+8EiGRiFKQ8e64Lh2NI4haZxMVVv/jw7PwWG27OZMenBJMecnkvyZetJfB1DvoJYOmPchmqx2AoU0aMigbYWyT2QFagGo4plcO1uRvZdiWu30laQrb3TNhV7xzeENYWD5y/GG94pDHKvKHr9ibbAcmq7dMyTLwXV7lKoRc0dRzQU30VtV53BFJwtlGuPan+eilnnvQsJ5ChSk5KvyvJ0geqBbM0JJWT5TIWL6YnFpNckWq2CY621oii8wbR0oeV2jFZkIzGfPd6Hz9difV5MnqDFqiw3m02xsl3hwqoc82KJK29OzotJ0aTWMIdEoY03y9mAdiAbN7haUSi0K3chJeukk+GQ0/dDIyCB7TDQnxTnxeQkqOLsLdf1LqYW7RHTxyU37qwn3fcH///HbTheWqK/U+kNasuoOwX60cj3sDcyr7cBtBr33VwCe5aD+p4h7oLJmT9/7yjwGM0lrDFoXHDLvFN15OcaqiWaSP/C/7fbcZX+Lg6r91W6e+PZLUuJpuM3kPBA28NqzvMs9xZmIsPhhVLk01Hai6WRj2f88ur66ssVSMBxAg+e53py/8AAr7J6PjMDC/7N8icpfT9MVM6P8cPRTzMeB3WIZpHmOecfJcJtbA==",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},B=void 0,m={},D=[];function J(e){const t={p:"p",...(0,a.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Delete a Role"}),"\n",(0,n.jsx)(l(),{method:"delete",path:"/user-svc/role/{roleId}",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Delete a role based on the role ID."}),"\n",(0,n.jsx)(b.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(d(),{parameters:[{description:"Role ID",in:"path",name:"roleId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function x(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(J,{...e})}):J(e)}}}]);