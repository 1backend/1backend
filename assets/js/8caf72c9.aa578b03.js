"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2482],{71385:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>D,contentTitle:()=>g,default:()=>m,frontMatter:()=>h,metadata:()=>a,toc:()=>f});const a=JSON.parse('{"id":"1backend/delete-user","title":"Delete a User","description":"Delete a user based on the user ID.","source":"@site/docs/1backend/delete-user.api.mdx","sourceDirName":"1backend","slug":"/1backend/delete-user","permalink":"/docs/1backend/delete-user","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"delete-user","title":"Delete a User","description":"Delete a user based on the user ID.","sidebar_label":"Delete a User","hide_title":true,"hide_table_of_contents":true,"api":"eJylVNuO2zYQ/RVinhqAK3l3G6DVUx3sIjCyQII6+7TxA02NJWYpkiEpu65AoB/RL+yXFCPKa+8lQZC8SJQ4lzPnzMwANQbplYvKGqjgCjVGZIL1AT1bi4A1s4bFFvOfxVUBHKxDL8hjUUMF9ehzG9ADBye86DCiD1DdDU+i3+YQwEHRpxOxBQ5GdAgVUPxFDRw8fumVxxqq6HvkEGSLnYBqgLh3ZBmiV6aBlFZkHJw1AQPdX8xm9Hqc9P074CCtiWgi3QrntJIj/PJzIJPheQq7/owyQkopcfh1dv487K0RfWytV39j/QMJHmqgBK9fwr0wEb0Rmi3Rb9Gza++t/7lMiUNA2XsV96M6b1B49PM+tlDdrYjOKBoSLiu13EpYcegwtvYo9CgyeUBJkp2FrRwP5ZAVTEBZCHJugd5rsi21lUK3NsTq9W8Xl+dA6Q5olgQ6a3iK6SklH/cO2afJ5BOwjdXa7rBm6z0TLDghkQlTs2jv0TAhcx+xjbfd2MKHqtiNbZRhaGpnlYnFoSFbFPXYxVNLzieFR4aBHxgVTr3DPRCdymws4SRRhBxFwU4oqjgIjeGPoEzTaxG9NYW03UnsDwu27J2zPgKfSGpjdFVZnq+FvEdTk0MJiT9hYc56ozaKys52rMYtaus6NJE5LeLG+o5trGd723s2X7CTRgn//fNvp6S3pJGSGM7ylBNv617pyKJlQQqNRItWEk1AquuA++2HG7a9LGaPUIeqLHe7XdGYvrC+KSe/UIrG6bPLYla0sdNUS0TfhfebZc5+LDrsRNOgL5QtR5OS+FZRk8n5m1wocKC2yjTMistiduZlcfE7xXU2xE6YE6QPu2xaTY9YHI5z9J1LbxI/4l+xdFooQ1lHBoZpIO7gMBDETU5aTWttxYF6n4yGgVLcep0S/f7So6dxXHHYCq/Emkqm1akCnWuoNkIH/Ab+X/6cNuYrdtywL8I9NLDZE5VC9/QFHO5xf9zAaZX4YRQISL6cS4kunrg9Wz7pdFdcXd9cf7wGDmKa5OPsUDx+OFCCF1E9nb2Mgp6Jf8VlGPJkpvRgn6++6vEw8NmaSFqllP4Ha6VobA==","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Create a New User","permalink":"/docs/1backend/create-user"},"next":{"title":"Save User","permalink":"/docs/1backend/save-user"}}');var n=s(74848),d=s(28453),i=s(53746),r=s.n(i),l=s(56518),o=s.n(l),c=s(99972),p=s.n(c),u=s(25342),x=s.n(u),b=(s(44215),s(82223),s(24861));const h={id:"delete-user",title:"Delete a User",description:"Delete a user based on the user ID.",sidebar_label:"Delete a User",hide_title:!0,hide_table_of_contents:!0,api:"eJylVNuO2zYQ/RVinhqAK3l3G6DVUx3sIjCyQII6+7TxA02NJWYpkiEpu65AoB/RL+yXFCPKa+8lQZC8SJQ4lzPnzMwANQbplYvKGqjgCjVGZIL1AT1bi4A1s4bFFvOfxVUBHKxDL8hjUUMF9ehzG9ADBye86DCiD1DdDU+i3+YQwEHRpxOxBQ5GdAgVUPxFDRw8fumVxxqq6HvkEGSLnYBqgLh3ZBmiV6aBlFZkHJw1AQPdX8xm9Hqc9P074CCtiWgi3QrntJIj/PJzIJPheQq7/owyQkopcfh1dv487K0RfWytV39j/QMJHmqgBK9fwr0wEb0Rmi3Rb9Gza++t/7lMiUNA2XsV96M6b1B49PM+tlDdrYjOKBoSLiu13EpYcegwtvYo9CgyeUBJkp2FrRwP5ZAVTEBZCHJugd5rsi21lUK3NsTq9W8Xl+dA6Q5olgQ6a3iK6SklH/cO2afJ5BOwjdXa7rBm6z0TLDghkQlTs2jv0TAhcx+xjbfd2MKHqtiNbZRhaGpnlYnFoSFbFPXYxVNLzieFR4aBHxgVTr3DPRCdymws4SRRhBxFwU4oqjgIjeGPoEzTaxG9NYW03UnsDwu27J2zPgKfSGpjdFVZnq+FvEdTk0MJiT9hYc56ozaKys52rMYtaus6NJE5LeLG+o5trGd723s2X7CTRgn//fNvp6S3pJGSGM7ylBNv617pyKJlQQqNRItWEk1AquuA++2HG7a9LGaPUIeqLHe7XdGYvrC+KSe/UIrG6bPLYla0sdNUS0TfhfebZc5+LDrsRNOgL5QtR5OS+FZRk8n5m1wocKC2yjTMistiduZlcfE7xXU2xE6YE6QPu2xaTY9YHI5z9J1LbxI/4l+xdFooQ1lHBoZpIO7gMBDETU5aTWttxYF6n4yGgVLcep0S/f7So6dxXHHYCq/Emkqm1akCnWuoNkIH/Ab+X/6cNuYrdtywL8I9NLDZE5VC9/QFHO5xf9zAaZX4YRQISL6cS4kunrg9Wz7pdFdcXd9cf7wGDmKa5OPsUDx+OFCCF1E9nb2Mgp6Jf8VlGPJkpvRgn6++6vEw8NmaSFqllP4Ha6VobA==",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},g=void 0,D={},f=[];function k(e){const t={p:"p",...(0,d.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Delete a User"}),"\n",(0,n.jsx)(r(),{method:"delete",path:"/user-svc/user/{userId}",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Delete a user based on the user ID."}),"\n",(0,n.jsx)(b.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(o(),{parameters:[{description:"User ID",in:"path",name:"userId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(x(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function m(e={}){const{wrapper:t}={...(0,d.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(k,{...e})}):k(e)}}}]);