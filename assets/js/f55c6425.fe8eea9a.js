"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3469],{24219:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>h,contentTitle:()=>j,default:()=>g,frontMatter:()=>k,metadata:()=>n,toc:()=>v});const n=JSON.parse('{"id":"1backend/save-user","title":"Save User","description":"Save user information based on the provided user ID.","source":"@site/docs/1backend/save-user.api.mdx","sourceDirName":"1backend","slug":"/1backend/save-user","permalink":"/docs/1backend/save-user","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"save-user","title":"Save User","description":"Save user information based on the provided user ID.","sidebar_label":"Save User","hide_title":true,"hide_table_of_contents":true,"api":"eJzdVk1z2zYQ/SuYPbUzNClX6bTDU+0m6ajJ1K4VTw+2plmRKxIxCCAAKFnl8L93lhAt2VIz7SWHXmwQ2I+3b7EP6qAkXzhpgzQacpjjmkTryQmpV8Y1yPtiiZ5KYbQINQnrzFqWVEaz2ev0Xs+CkF5IHUjzwco4gWUjtU/EkgpsPQkZ2N4PET6y55lfFzkvcipl+CgsuUZ6z+k2tSxqYbTa7sKIGteU3uu3HDjmDUawn5DBC7PRjGolFSXCEwmPa5qTWqWQgLHkhipmJeTAJ7eeHCRg0WFDgZyH/K57wcNtrA0SkPxpMdSQgMaGIAcGMCshAUefW+mohDy4lhLwRU0NQt5B2Fq29MFJXUHfL6Ix+XBpyi1bFIbpCrxEa5UsBpDZJ8/5u4NQ1nEJQZLnrwjhKEECXrXVyYNQt81So1RvpSImoQN6xMYqNmPO/lz9/vqx9X/8aCpIjpA/7ZjlJyoC9Lx14tJcxwaIm1jmETs9b3hrtI+FfDeZ8L/nka7eQfLvmTkC1ifw6lTYmV6jkqX4dX71239J8Jx6cs64U809QdGA5PwYya3GNtTGyb+o/GpIvj/NSSCnUYk5uTU58WaI+XUg8XWlonUybIfZuyR05C7aUEN+t+BhCVjxWMY5nK8LWCTQUKgND7FtwzC/bA7ZKCbDIuvicPbAKbiwON2tU2ybKVOgqo0P+fn5dPoDcK4RypxLi9UcAnpJ3IetJXG/M7kHsTJKmQ2VYrkVKLzFggTqUgTzQFpgEYdArJxpBvUbSxLvTSW1IF1aI3VIR62pCctBoHZqc7G7MEMf9gOKVr6j7UAw9+Zmry5vxuke1WJsxigS++8jbTgShJ5RrcyoWFgMN4MalGqQU0X+Jy911SoMzui0MM0B9OuZmLfWGscNiz2oQ7B5lp0vsXggXbJDBkeScjE70xjkmkQjC2e4lbIgL6zCwO8Ss6VkQdof1nnxy/V7sZ6mk2fZfJ5lm80mrXSbGldlOz+fYWXV2TSdpHVo1CCV5Bp/tZrHbHuwfoNVRS6VJhtMMqZOhkFBzy9jIZAA37YIf5JO08mZK9LpK45rjQ8N6gOkg2TuHqJnlR88DP/Dx3h3dwM9hswqlJrZGTrV7eb5DkY83MNIUL57cBcJ8OiyUdcxDbdO9T1vf27JsZQsElijk7jk1vCjLj2vS8hXqDx9getvbnav1bdi//afhDvOn95yy1G1/AUJPNB2/9ugX/TJOMkMJB7+HNOdsYQcOB/pbJ+MHhdFQTZ80fZQGa9vP0ACy91PjMaU7OJww68xbiJOM1Q/yNyw14FCXbVYsW0MyaKCO/Hbyw0jSsYFF3WSiZdyFevgv1zVSZeui2LW90/28egfPZ40MlpzYxZ93/8Nen65SA==","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Delete a User","permalink":"/docs/1backend/delete-user"},"next":{"title":"Read User by Token","permalink":"/docs/1backend/read-user-by-token"}}');var a=t(74848),r=t(28453),i=t(53746),o=t.n(i),p=t(56518),d=t.n(p),c=t(99972),l=t.n(c),b=t(25342),u=t.n(b),m=(t(44215),t(82223),t(24861));const k={id:"save-user",title:"Save User",description:"Save user information based on the provided user ID.",sidebar_label:"Save User",hide_title:!0,hide_table_of_contents:!0,api:"eJzdVk1z2zYQ/SuYPbUzNClX6bTDU+0m6ajJ1K4VTw+2plmRKxIxCCAAKFnl8L93lhAt2VIz7SWHXmwQ2I+3b7EP6qAkXzhpgzQacpjjmkTryQmpV8Y1yPtiiZ5KYbQINQnrzFqWVEaz2ev0Xs+CkF5IHUjzwco4gWUjtU/EkgpsPQkZ2N4PET6y55lfFzkvcipl+CgsuUZ6z+k2tSxqYbTa7sKIGteU3uu3HDjmDUawn5DBC7PRjGolFSXCEwmPa5qTWqWQgLHkhipmJeTAJ7eeHCRg0WFDgZyH/K57wcNtrA0SkPxpMdSQgMaGIAcGMCshAUefW+mohDy4lhLwRU0NQt5B2Fq29MFJXUHfL6Ix+XBpyi1bFIbpCrxEa5UsBpDZJ8/5u4NQ1nEJQZLnrwjhKEECXrXVyYNQt81So1RvpSImoQN6xMYqNmPO/lz9/vqx9X/8aCpIjpA/7ZjlJyoC9Lx14tJcxwaIm1jmETs9b3hrtI+FfDeZ8L/nka7eQfLvmTkC1ifw6lTYmV6jkqX4dX71239J8Jx6cs64U809QdGA5PwYya3GNtTGyb+o/GpIvj/NSSCnUYk5uTU58WaI+XUg8XWlonUybIfZuyR05C7aUEN+t+BhCVjxWMY5nK8LWCTQUKgND7FtwzC/bA7ZKCbDIuvicPbAKbiwON2tU2ybKVOgqo0P+fn5dPoDcK4RypxLi9UcAnpJ3IetJXG/M7kHsTJKmQ2VYrkVKLzFggTqUgTzQFpgEYdArJxpBvUbSxLvTSW1IF1aI3VIR62pCctBoHZqc7G7MEMf9gOKVr6j7UAw9+Zmry5vxuke1WJsxigS++8jbTgShJ5RrcyoWFgMN4MalGqQU0X+Jy911SoMzui0MM0B9OuZmLfWGscNiz2oQ7B5lp0vsXggXbJDBkeScjE70xjkmkQjC2e4lbIgL6zCwO8Ss6VkQdof1nnxy/V7sZ6mk2fZfJ5lm80mrXSbGldlOz+fYWXV2TSdpHVo1CCV5Bp/tZrHbHuwfoNVRS6VJhtMMqZOhkFBzy9jIZAA37YIf5JO08mZK9LpK45rjQ8N6gOkg2TuHqJnlR88DP/Dx3h3dwM9hswqlJrZGTrV7eb5DkY83MNIUL57cBcJ8OiyUdcxDbdO9T1vf27JsZQsElijk7jk1vCjLj2vS8hXqDx9getvbnav1bdi//afhDvOn95yy1G1/AUJPNB2/9ugX/TJOMkMJB7+HNOdsYQcOB/pbJ+MHhdFQTZ80fZQGa9vP0ACy91PjMaU7OJww68xbiJOM1Q/yNyw14FCXbVYsW0MyaKCO/Hbyw0jSsYFF3WSiZdyFevgv1zVSZeui2LW90/28egfPZ40MlpzYxZ93/8Nen65SA==",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},j=void 0,h={},v=[];function y(e){const s={code:"code",p:"p",...(0,r.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Save User"}),"\n",(0,a.jsx)(o(),{method:"put",path:"/user-svc/user/{userId}",context:"endpoint"}),"\n",(0,a.jsxs)(s.p,{children:["Save user information based on the provided user ID.\nIt is intended for admins, because it uses the ",(0,a.jsx)(s.code,{children:"user-svc:user:edit"})," permission which only admins have.\nFor a user to edit its own profile, see saveSelf."]}),"\n",(0,a.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,a.jsx)(d(),{parameters:[{description:"User ID",in:"path",name:"userId",required:!0,schema:{type:"string"}}]}),"\n",(0,a.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{name:{type:"string"},slug:{type:"string"},thumbnailFileId:{example:"file_fQDxusW8og",type:"string"}},type:"object"}}},description:"Save Profile Request",required:!0}}),"\n",(0,a.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function g(e={}){const{wrapper:s}={...(0,r.R)(),...e.components};return s?(0,a.jsx)(s,{...e,children:(0,a.jsx)(y,{...e})}):y(e)}}}]);