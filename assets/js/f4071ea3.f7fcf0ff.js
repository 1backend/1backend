"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3488],{21445:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>v,contentTitle:()=>f,default:()=>Q,frontMatter:()=>h,metadata:()=>n,toc:()=>x});const n=JSON.parse('{"id":"1backend/subscribe-to-events","title":"Subscribe to the Event Stream","description":"Establish a subscription to the firehose events and accept a real time stream of them.","source":"@site/docs/1backend/subscribe-to-events.api.mdx","sourceDirName":"1backend","slug":"/1backend/subscribe-to-events","permalink":"/docs/1backend/subscribe-to-events","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"subscribe-to-events","title":"Subscribe to the Event Stream","description":"Establish a subscription to the firehose events and accept a real time stream of them.","sidebar_label":"Subscribe to the Event Stream","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VE2P2zYQ/SvEnGnJGzdAoVM3wCZYNEACePfk+DCmxhKzFMmQlB1X4H8vRrK8jpscemgvlkHOx5s373GAmqIK2iftLFTwEBPujI6tQBH73eVKJCdSS2KvA7UukqAD2RQF2lqgUuSTQBEIjUi6IxFTIOyE23NSV4AE5ykgV3qsoYJz6R09uYexEEgIFL2zkSJUA7xZLvlzA45DRY0JQYJyNpFNHJXoeypHQIupMR9G1VKH4/XJE/dMQdsGcs5Zwm/Lu3/Wf7bYp9YF/RfV/6qDDzxe0hN2CsGFnzWW84nbfSWVzlDe/mzUR5soWDRiTeFAQTyMNf8nTFlCJNUHnU5QbQZ4Rxgo3PephWqzzVsJCZsI1Qbez3JYHxRsJXSUWscLbiiBBI+cAuWsmkU8qAlyLC8SAO7GQ8axWR8Mp5TGKTSti6l6+/ub1R1w2xnVmqecBrvGdkvi08mT+HIO+QJi74xxR6rF7sTy9qhoFHByL2QFqm+9DlSLfXDdKPbnSIEnEx9do60gW3unbWI5a67fEtbEW7HYMYX3Z/WMOocLs+j1n3QaudZ27xgnrxHVuEbqUPPEEQ3FP6K2TW8wBWcL5bqr2p8fxbr33gUmdiKpTclXZXm3Q/VCtuaEErK8YeH+cWEx6QOJTqvgmGutKApvMO1dGN1ptCIbifHM/T58/igOq2L5Q7dYleXxeCwa2xcuNOU5L5bYeLNYFcuiTZ1hDIlCFz/t11O3V7DxiE1DodCuHENK5kknwyF376ZBQALLYYK/LFbFchFUsVpxXe9i6tBeIV3PQprfqOmZWE+euGFjeHXQf/fWnfc+GtQb1JaBjyQOZ0ts4NoSIIHmV/DVFlsJLH8OHoYdRnoOJmc+/tZTYGduJRwwaNwxe5ttlrMi2UcvdOI1jniZUDT9BdQPr0a+9u2HhyeQgGc7vQqYq8n5D5efr+zpqvitASYM/JvlL1KGYbJHzpf46eqXGRfXTdHM7jbn/DdgRFXT","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Publish an Event","permalink":"/docs/1backend/publish-event"},"next":{"title":"Start the Default Model","permalink":"/docs/1backend/start-default-model"}}');var i=s(74848),r=s(28453),a=s(53746),o=s.n(a),c=s(56518),d=s.n(c),b=s(99972),l=s.n(b),u=s(25342),p=s.n(u),m=(s(44215),s(82223),s(24861));const h={id:"subscribe-to-events",title:"Subscribe to the Event Stream",description:"Establish a subscription to the firehose events and accept a real time stream of them.",sidebar_label:"Subscribe to the Event Stream",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VE2P2zYQ/SvEnGnJGzdAoVM3wCZYNEACePfk+DCmxhKzFMmQlB1X4H8vRrK8jpscemgvlkHOx5s373GAmqIK2iftLFTwEBPujI6tQBH73eVKJCdSS2KvA7UukqAD2RQF2lqgUuSTQBEIjUi6IxFTIOyE23NSV4AE5ykgV3qsoYJz6R09uYexEEgIFL2zkSJUA7xZLvlzA45DRY0JQYJyNpFNHJXoeypHQIupMR9G1VKH4/XJE/dMQdsGcs5Zwm/Lu3/Wf7bYp9YF/RfV/6qDDzxe0hN2CsGFnzWW84nbfSWVzlDe/mzUR5soWDRiTeFAQTyMNf8nTFlCJNUHnU5QbQZ4Rxgo3PephWqzzVsJCZsI1Qbez3JYHxRsJXSUWscLbiiBBI+cAuWsmkU8qAlyLC8SAO7GQ8axWR8Mp5TGKTSti6l6+/ub1R1w2xnVmqecBrvGdkvi08mT+HIO+QJi74xxR6rF7sTy9qhoFHByL2QFqm+9DlSLfXDdKPbnSIEnEx9do60gW3unbWI5a67fEtbEW7HYMYX3Z/WMOocLs+j1n3QaudZ27xgnrxHVuEbqUPPEEQ3FP6K2TW8wBWcL5bqr2p8fxbr33gUmdiKpTclXZXm3Q/VCtuaEErK8YeH+cWEx6QOJTqvgmGutKApvMO1dGN1ptCIbifHM/T58/igOq2L5Q7dYleXxeCwa2xcuNOU5L5bYeLNYFcuiTZ1hDIlCFz/t11O3V7DxiE1DodCuHENK5kknwyF376ZBQALLYYK/LFbFchFUsVpxXe9i6tBeIV3PQprfqOmZWE+euGFjeHXQf/fWnfc+GtQb1JaBjyQOZ0ts4NoSIIHmV/DVFlsJLH8OHoYdRnoOJmc+/tZTYGduJRwwaNwxe5ttlrMi2UcvdOI1jniZUDT9BdQPr0a+9u2HhyeQgGc7vQqYq8n5D5efr+zpqvitASYM/JvlL1KGYbJHzpf46eqXGRfXTdHM7jbn/DdgRFXT",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,v={},x=[];function E(e){const t={p:"p",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Subscribe to the Event Stream"}),"\n",(0,i.jsx)(o(),{method:"get",path:"/firehose-svc/events/subscribe",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Establish a subscription to the firehose events and accept a real time stream of them."}),"\n",(0,i.jsx)(d(),{parameters:void 0}),"\n",(0,i.jsx)(l(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(p(),{id:void 0,label:void 0,responses:{200:{description:"Event data",content:{"text/event-stream":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"text/event-stream":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"text/event-stream":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function Q(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(E,{...e})}):E(e)}}}]);