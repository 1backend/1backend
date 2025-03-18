"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7327],{67995:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>D,contentTitle:()=>u,default:()=>m,frontMatter:()=>j,metadata:()=>i,toc:()=>x});const i=JSON.parse('{"id":"1backend/delete-thread","title":"Delete a Thread","description":"Delete a specific chat thread by its ID","source":"@site/docs/1backend/delete-thread.api.mdx","sourceDirName":"1backend","slug":"/1backend/delete-thread","permalink":"/docs/1backend/delete-thread","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"delete-thread","title":"Delete a Thread","description":"Delete a specific chat thread by its ID","sidebar_label":"Delete a Thread","hide_title":true,"hide_table_of_contents":true,"api":"eJylVdtu4zYQ/RVinlpAkZykC7R6Wm8TFG6DblBvnrJ+GFMjixuK5JKUXVcg0I/oF/ZLFiP5tom3QNAXX8QZnjOH51A9VBSkVy4qa6CEG9IUSaAIjqSqlRSywShi4wkrsdwKFYOY3UAG1pFH7ppVUEI19H0YyiADhx5biuQDlI/9M4yxatxF8QOHsYEMDLYEJYxYM97G0+dOeaqgjL6jDIJsqEUoe4hbx7UhemVWkNKCi4OzJlDg9avJhL/OAodOSgqh7rTeipE4g0lrIpnIbeicVnKYrvgUuLc/wcaqUryE+t6zClEx5shwx8suP5GMkFJKGfxwjsvMrFGrSvw6f//7a8CfDz4CXL4EeDDYxcZ69dfrpjsH8Ob8BJG8QS3m5Nfkxa331v8/pJRBINl5FbeDb94RevLTLjZQPi74kCOu2FLwM7tyvpawyKCl2NijCQf7cQcU7N2LsJbF6Kmi33srASMx7dGgnddcX2grUTc2xPLNj1fXl8CQe0ZzJj6665TXC5NtHYmPu5KPIGqrtd3QEB6OFUoSaCoR7RMZgXJ0uKi9bUVsSDwE8jyZuLMrZQSZylllYr4PS0NYkT/GZbo75UFlODgQnfqNtsCSKlNb5skHg3I4GGpR8cQBNYW3QZlVpzF6a3Jp25O972di3jlnfYRsJ1IToyuL4nKJ8olMxQ0FpOyZClPRGVUrHnusExWtSVvXkonCaYy19a2orRdb23kxnYkTs4R///6nVdJbPiMlKVwsMVA16LbslI4iWhEkamJZtJJkAvFce96/3N+J9XU++Yp1KItis9nkK9Pl1q+KXV8ocOX0xXU+yZvYap4lkm/D+3o+oh+HDhtcrcjnyhZDScF6q6i55PLdOChkwLYaZZjk1/nkwsv86ife19kQWzQnTA/37eHq/ErH/pimV1zNOwtE+jMWTqMyjD3o0O+i8Qj7aHD5Hro8XL2LDDgFXNj3rP2D1ynx488deQ7nIoM1eoVLHp6veBX4dwVljTrQf8zx3R+7W/17cfomOEt6b2azZVlRd/wPMnii7embIi1Stg8GkxmXp1KSiyeNL66jdHp73Nze3X64hQxwl+tjkni/bP+DAc7yep7EkQV/puwbLX0/5jSlQ/249M2OQ/zHapZpkVL6AmU2qMI=","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Add Thread","permalink":"/docs/1backend/add-thread"},"next":{"title":"Get Thread","permalink":"/docs/1backend/get-thread"}}');var n=a(74848),d=a(28453),s=a(53746),r=a.n(s),l=a(56518),c=a.n(l),o=a(99972),p=a.n(o),h=a(25342),k=a.n(h),b=(a(44215),a(82223),a(24861));const j={id:"delete-thread",title:"Delete a Thread",description:"Delete a specific chat thread by its ID",sidebar_label:"Delete a Thread",hide_title:!0,hide_table_of_contents:!0,api:"eJylVdtu4zYQ/RVinlpAkZykC7R6Wm8TFG6DblBvnrJ+GFMjixuK5JKUXVcg0I/oF/ZLFiP5tom3QNAXX8QZnjOH51A9VBSkVy4qa6CEG9IUSaAIjqSqlRSywShi4wkrsdwKFYOY3UAG1pFH7ppVUEI19H0YyiADhx5biuQDlI/9M4yxatxF8QOHsYEMDLYEJYxYM97G0+dOeaqgjL6jDIJsqEUoe4hbx7UhemVWkNKCi4OzJlDg9avJhL/OAodOSgqh7rTeipE4g0lrIpnIbeicVnKYrvgUuLc/wcaqUryE+t6zClEx5shwx8suP5GMkFJKGfxwjsvMrFGrSvw6f//7a8CfDz4CXL4EeDDYxcZ69dfrpjsH8Ob8BJG8QS3m5Nfkxa331v8/pJRBINl5FbeDb94RevLTLjZQPi74kCOu2FLwM7tyvpawyKCl2NijCQf7cQcU7N2LsJbF6Kmi33srASMx7dGgnddcX2grUTc2xPLNj1fXl8CQe0ZzJj6665TXC5NtHYmPu5KPIGqrtd3QEB6OFUoSaCoR7RMZgXJ0uKi9bUVsSDwE8jyZuLMrZQSZylllYr4PS0NYkT/GZbo75UFlODgQnfqNtsCSKlNb5skHg3I4GGpR8cQBNYW3QZlVpzF6a3Jp25O972di3jlnfYRsJ1IToyuL4nKJ8olMxQ0FpOyZClPRGVUrHnusExWtSVvXkonCaYy19a2orRdb23kxnYkTs4R///6nVdJbPiMlKVwsMVA16LbslI4iWhEkamJZtJJkAvFce96/3N+J9XU++Yp1KItis9nkK9Pl1q+KXV8ocOX0xXU+yZvYap4lkm/D+3o+oh+HDhtcrcjnyhZDScF6q6i55PLdOChkwLYaZZjk1/nkwsv86ife19kQWzQnTA/37eHq/ErH/pimV1zNOwtE+jMWTqMyjD3o0O+i8Qj7aHD5Hro8XL2LDDgFXNj3rP2D1ynx488deQ7nIoM1eoVLHp6veBX4dwVljTrQf8zx3R+7W/17cfomOEt6b2azZVlRd/wPMnii7embIi1Stg8GkxmXp1KSiyeNL66jdHp73Nze3X64hQxwl+tjkni/bP+DAc7yep7EkQV/puwbLX0/5jSlQ/249M2OQ/zHapZpkVL6AmU2qMI=",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},u=void 0,D={},x=[];function y(e){const t={p:"p",...(0,d.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Delete a Thread"}),"\n",(0,n.jsx)(r(),{method:"delete",path:"/chat-svc/thread/{threadId}",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Delete a specific chat thread by its ID"}),"\n",(0,n.jsx)(b.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(c(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(k(),{id:void 0,label:void 0,responses:{200:{description:"Thread successfully deleted",content:{"application/json":{schema:{additionalProperties:!0,type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function m(e={}){const{wrapper:t}={...(0,d.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(y,{...e})}):y(e)}}}]);