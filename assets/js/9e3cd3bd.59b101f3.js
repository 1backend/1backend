"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[702],{63105:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>k,contentTitle:()=>v,default:()=>y,frontMatter:()=>m,metadata:()=>s,toc:()=>C});const s=JSON.parse('{"id":"1backend/events","title":"Events","description":"Events is a dummy endpoint to display documentation about the events that this service emits.","source":"@site/docs/1backend/events.api.mdx","sourceDirName":"1backend","slug":"/1backend/events","permalink":"/docs/1backend/events","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"events","title":"Events","description":"Events is a dummy endpoint to display documentation about the events that this service emits.","sidebar_label":"Events","hide_title":true,"hide_table_of_contents":true,"api":"eJy1U8tu2zAQ/BViz7Jk14cWOrUpgiBIgARwcnJ8WFNriQlFMuTKriro34u17Lzu1UUSuY+Z2dkBKko6msDGOyjhck+OkzJJoaq6tu0VuSp441ixV5VJwWKvKq+7lhyjZCnc+o4VN6RoyuYG5d8klSjujSZFreGUQwY+UDxmXVdQwhQPGURKwbtECcoBvs3n8voM7O4GMtDeMTmWWwzBGn0sVTwnCRkg6YZalK8QpRGbqSA3kbCSjgNwHwhKSByNq2Ecs/OJ3z6TZhjlyYCxTlCu4bdQWe01bDJoiRsvsGtiyCAgN1BCoRvkWdrr4o2NsKYo+QN00UpQYb1G2/jE5WKxXH6HcSNxuouG+5XgnqBeEEaKvzop/VWChz6QejqFPIHaeWv9gSq17RWqFFCTQlcp9i/kFOrXzkSq1C769jidx0RRuKhbXxv3NlgZi5H6DWFFETJw2IoigsJH8/coMrwJhcHcUH+UzridF5wyF9THuVCLRhgntJR+JuPqziJH73Lt2w+176/VqgvBR5FyEqlhDmVRLLaoX8hVklDAmH1R4df1zCGbPanW6OhPDksqWOSdj63QsUaTSyR4zv2u7m/VfpnPP3VLZVEcDoe8dl3uY12c8lKBdbCzZT7PG26tYGCKbbrbraZu72DTAeuaYm58cQwpRCfDVkIWFxMRyEDsMMGf58t8Pos6X/6QusEnbtF9QHp5dtEn2sO79//7jp7mzPSHi2DROAF6FG04mX4NZ9NDdl7iTQZib7kchi0meox2HOX4taPYQ7neZLDHaHAr6qw3Y3Z2nOzJC/UyJq0piCP2aLvJbF/WfPy4iFeXDzCO/wBPc6NV","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Introduction","permalink":"/docs/1backend/1-backend"},"next":{"title":"Delete a Message","permalink":"/docs/1backend/delete-message"}}');var i=n(74848),a=n(28453),o=n(53746),d=n.n(o),c=n(56518),r=n.n(c),u=n(99972),l=n.n(u),p=n(25342),b=n.n(p),h=(n(44215),n(82223),n(24861));const m={id:"events",title:"Events",description:"Events is a dummy endpoint to display documentation about the events that this service emits.",sidebar_label:"Events",hide_title:!0,hide_table_of_contents:!0,api:"eJy1U8tu2zAQ/BViz7Jk14cWOrUpgiBIgARwcnJ8WFNriQlFMuTKriro34u17Lzu1UUSuY+Z2dkBKko6msDGOyjhck+OkzJJoaq6tu0VuSp441ixV5VJwWKvKq+7lhyjZCnc+o4VN6RoyuYG5d8klSjujSZFreGUQwY+UDxmXVdQwhQPGURKwbtECcoBvs3n8voM7O4GMtDeMTmWWwzBGn0sVTwnCRkg6YZalK8QpRGbqSA3kbCSjgNwHwhKSByNq2Ecs/OJ3z6TZhjlyYCxTlCu4bdQWe01bDJoiRsvsGtiyCAgN1BCoRvkWdrr4o2NsKYo+QN00UpQYb1G2/jE5WKxXH6HcSNxuouG+5XgnqBeEEaKvzop/VWChz6QejqFPIHaeWv9gSq17RWqFFCTQlcp9i/kFOrXzkSq1C769jidx0RRuKhbXxv3NlgZi5H6DWFFETJw2IoigsJH8/coMrwJhcHcUH+UzridF5wyF9THuVCLRhgntJR+JuPqziJH73Lt2w+176/VqgvBR5FyEqlhDmVRLLaoX8hVklDAmH1R4df1zCGbPanW6OhPDksqWOSdj63QsUaTSyR4zv2u7m/VfpnPP3VLZVEcDoe8dl3uY12c8lKBdbCzZT7PG26tYGCKbbrbraZu72DTAeuaYm58cQwpRCfDVkIWFxMRyEDsMMGf58t8Pos6X/6QusEnbtF9QHp5dtEn2sO79//7jp7mzPSHi2DROAF6FG04mX4NZ9NDdl7iTQZib7kchi0meox2HOX4taPYQ7neZLDHaHAr6qw3Y3Z2nOzJC/UyJq0piCP2aLvJbF/WfPy4iFeXDzCO/wBPc6NV",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},v=void 0,k={},C=[];function R(e){const t={p:"p",...(0,a.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Events"}),"\n",(0,i.jsx)(d(),{method:"get",path:"/chat-svc/events",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Events is a dummy endpoint to display documentation about the events that this service emits."}),"\n",(0,i.jsx)(r(),{parameters:void 0}),"\n",(0,i.jsx)(l(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{threadId:{type:"string"}},type:"object"}}}}}})]})}function y(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(R,{...e})}):R(e)}}}]);