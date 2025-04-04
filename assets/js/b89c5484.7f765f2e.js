"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2834],{83594:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>g,contentTitle:()=>v,default:()=>O,frontMatter:()=>u,metadata:()=>s,toc:()=>k});const s=JSON.parse('{"id":"1backend/remove-instance","title":"Remove Instance","description":"Removes a registered instance by ID.","source":"@site/docs/1backend/remove-instance.api.mdx","sourceDirName":"1backend","slug":"/1backend/remove-instance","permalink":"/docs/1backend/remove-instance","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"remove-instance","title":"Remove Instance","description":"Removes a registered instance by ID.","sidebar_label":"Remove Instance","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VU1v2zgQ/SvEnLqAIjl1CxQ6bdoEC2OD3SJpTqkPNDWW2FAkO6TsugL/+2IkObETd9FTLv6Q3nDeG75H9lBhUKR91M5CCTfYug0GIQVhrUNEwkpoG6K0CsVqJxaXOWTgPJLkkkUFJdBQtJhQkIGXJFuMSAHK+/5Ziz1OLC4hA82PvIwNZGBli1CCriADwu+dJqygjNRhBkE12Eooe4g7z6gQSdsaUloyOHhnAwZ+/3b2jr+Om/7jxCdnI9oIKYN3s9lLyMJupNHVSEtN4LIH6b3RalBbfAsM7Q/YeOJZRD32RiJHp0hm+ydu9Q1VhJTSwOP8JY87K7vYONI/sXpFJieGdou00QqFdVGsXWdfj8770xsUkaw0gnkhiathzdehlDIIqDrScTdY+iNKQrroYgPl/ZI9GGXNboebITe0E7cbBcsMWoyN45RUaDCO6eAqKGhCnoWNKvYZK3pdJeBurHHMT0eG8YVxSprGhVi+//B2fg7cds/qllWOwg65PZ/hl51H8XWCfAWxdsa4LVYcbSmClwqFtJWI7gGtkGoMoViTa0VsUNwFJFYmrl2trUBbeadtzPdJblBWSE9ZvpisPGwJPE5Wev037oZZa7t2zJN3UaphF7GVmhUHaTD8GbStOyMjOZsr1x6s/XkhbjvvHUXIpiE1MfqyKM5XUj2grbig4MQfT+FicWZl1BsUrVbkwujzILyRce2oZTlGK7QBmc++31+fr8Vmns+OuoWyKLbbbV7bLndUF1NdKGTtzdk8n+VNbA1ziEht+Hc9peqJbNjKukbKtSsGSMFz0tEw5PzjKAQyYDuM9Gf5PJ+dkcrnc17XuxBbaQ+Yjqe4ODiRj/T3T5H53QN/2reIP2LhjdSWGw9D6Cc/38Ohnwc/PHYvdcVRYOsysO9XMuAdmZT48fcOiVO1zGAjScsVK+drQwf+XUG5libg/4h4czPdFn+I49vlJO29B+2OpypNx/8ggwfcjbdPWqZs72QmMr64UAp9PCh5cdikw7hfXl1ffbmCDOQUxCfr83rZ/gc3OMnoeXRGFvyZsl+U9P0YrJQe8eOrX1Y85nUSzyJSSv8BsvjDBQ==","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Register Instance","permalink":"/docs/1backend/register-instance"},"next":{"title":"List Service Instances","permalink":"/docs/1backend/list-instances"}}');var i=t(74848),r=t(28453),a=t(53746),c=t.n(a),d=t(56518),o=t.n(d),l=t(99972),p=t.n(l),b=t(25342),h=t.n(b),m=(t(44215),t(82223),t(24861));const u={id:"remove-instance",title:"Remove Instance",description:"Removes a registered instance by ID.",sidebar_label:"Remove Instance",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU1v2zgQ/SvEnLqAIjl1CxQ6bdoEC2OD3SJpTqkPNDWW2FAkO6TsugL/+2IkObETd9FTLv6Q3nDeG75H9lBhUKR91M5CCTfYug0GIQVhrUNEwkpoG6K0CsVqJxaXOWTgPJLkkkUFJdBQtJhQkIGXJFuMSAHK+/5Ziz1OLC4hA82PvIwNZGBli1CCriADwu+dJqygjNRhBkE12Eooe4g7z6gQSdsaUloyOHhnAwZ+/3b2jr+Om/7jxCdnI9oIKYN3s9lLyMJupNHVSEtN4LIH6b3RalBbfAsM7Q/YeOJZRD32RiJHp0hm+ydu9Q1VhJTSwOP8JY87K7vYONI/sXpFJieGdou00QqFdVGsXWdfj8770xsUkaw0gnkhiathzdehlDIIqDrScTdY+iNKQrroYgPl/ZI9GGXNboebITe0E7cbBcsMWoyN45RUaDCO6eAqKGhCnoWNKvYZK3pdJeBurHHMT0eG8YVxSprGhVi+//B2fg7cds/qllWOwg65PZ/hl51H8XWCfAWxdsa4LVYcbSmClwqFtJWI7gGtkGoMoViTa0VsUNwFJFYmrl2trUBbeadtzPdJblBWSE9ZvpisPGwJPE5Wev037oZZa7t2zJN3UaphF7GVmhUHaTD8GbStOyMjOZsr1x6s/XkhbjvvHUXIpiE1MfqyKM5XUj2grbig4MQfT+FicWZl1BsUrVbkwujzILyRce2oZTlGK7QBmc++31+fr8Vmns+OuoWyKLbbbV7bLndUF1NdKGTtzdk8n+VNbA1ziEht+Hc9peqJbNjKukbKtSsGSMFz0tEw5PzjKAQyYDuM9Gf5PJ+dkcrnc17XuxBbaQ+Yjqe4ODiRj/T3T5H53QN/2reIP2LhjdSWGw9D6Cc/38Ohnwc/PHYvdcVRYOsysO9XMuAdmZT48fcOiVO1zGAjScsVK+drQwf+XUG5libg/4h4czPdFn+I49vlJO29B+2OpypNx/8ggwfcjbdPWqZs72QmMr64UAp9PCh5cdikw7hfXl1ffbmCDOQUxCfr83rZ/gc3OMnoeXRGFvyZsl+U9P0YrJQe8eOrX1Y85nUSzyJSSv8BsvjDBQ==",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},v=void 0,g={},k=[];function I(e){const n={p:"p",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Remove Instance"}),"\n",(0,i.jsx)(c(),{method:"delete",path:"/registry-svc/instance/{id}",context:"endpoint"}),"\n",(0,i.jsx)(n.p,{children:"Removes a registered instance by ID."}),"\n",(0,i.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(o(),{parameters:[{description:"Instance ID",in:"path",name:"id",required:!0,schema:{type:"string"}}]}),"\n",(0,i.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(h(),{id:void 0,label:void 0,responses:{204:{description:"No Content"},400:{description:"Invalid ID",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},404:{description:"Service not found",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function O(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(I,{...e})}):I(e)}}}]);