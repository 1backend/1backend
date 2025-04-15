"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[516],{76268:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>k,contentTitle:()=>R,default:()=>g,frontMatter:()=>u,metadata:()=>r,toc:()=>S});const r=JSON.parse('{"id":"1backend/subscribe-to-prompt-responses","title":"Subscribe to Prompt Responses by Thread","description":"Subscribe to prompt responses by thread via Server-Sent Events (SSE).","source":"@site/docs/1backend/subscribe-to-prompt-responses.api.mdx","sourceDirName":"1backend","slug":"/1backend/subscribe-to-prompt-responses","permalink":"/docs/1backend/subscribe-to-prompt-responses","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"subscribe-to-prompt-responses","title":"Subscribe to Prompt Responses by Thread","description":"Subscribe to prompt responses by thread via Server-Sent Events (SSE).","sidebar_label":"Subscribe to Prompt Responses by Thread","hide_title":true,"hide_table_of_contents":true,"api":"eJzdVU1v4zYQ/SvEnHYXtGTXBQro1GwbLIxu0SByDkViYGlqLLGRSC5J2XUF/fdiREt2vEl76qUXm5Lm483Me8MOCvTSKRuU0ZBB3m7pcYssGGadaWxgDr012qNn2yMLlUNRsL0SLEe3RzfLUQd2u0cdPHuX57fvkyf9u2mZFJr5y3DR1bMt7oxDFio8MuGQSYciYJE86XWFzAeHosGCDkqXfjAxOxaOFtmXfPj6U9Xq5y+ceRzCsLsB6Ppo0TPUhTVKB7YzjjWUqMAgVO0T4GAsOkGlrgrIYEK3NjHC/VgpcLDCiQYDOg/ZY3fVpnXswupn4KDohRWhAg5aNAgZxEpXBXBw+LVVDgvIgmuRg5cVNgKyDqggAjGUCX2/IeMxf9bBd/M5/V3NZ6hf6XKaCnCQRgfUgaw/pB/o7+0sfc/h+9ci/6q8p7gjdjY14F8yWEddDSqiRueMey0xH9+Y7R8owwRl8S2UBy3aUBmn/sLiv0vec/AoW6fCcZjwRxQO3U0bKsgeNzSOIEoaPkRysHwvYcOhwVAZok+JYeAJOUAaxTLze3k6+rQbe9mn02TTiXRA+UlBkWCtqylMWhsp6sr4kC0Wy+UPQEBGnDnVHUu9RPsNO0kqTyeTJ2A7U9fmgAXpVzBvhUQmdMGCeUbNhIwMZTtnmkFODx4dVcs+m1LpSVHJSPYKRTHQ4kT3m9O0BmXB1Gth1S94HLqv9M4QThqlkMMosRGKKvaiRv8jUa+tRXBGJ9I0F7HvVixvrTWOmh2bVIVgszRdbIV8Rl2QQwo9v+rCzWqmRVB7ZI2SzlCvlUTPbC3CzriGyqmVRNJQ1k35Pt19ZvtlMn+RzWdpejgcklK3iXFlevLzqShtPVsm86QKTU0YArrG/7bLY7YzWH8QZYkuUSYdTFLqkwo1mSw+xkKAA9Ehwp8ny2Q+czIhDnCwxodG6AukL1b1iaL3l6s6Lim46kt31tP/fdufeBjwz5DaWihNjRyG2p1k+whn2ZKWo3CBQ/ZihZ9vhbN6NxxIpRSj67bC44Or+55ef23R0UrZcNgLp8SWhkxXiPJ0LiDbidrjP8zl3f3p1njPLm+aV+sZxaaPRB9Rt/QEHJ7xeHkT9Zuej8IlMPHzjZRow4UjLdj+csl9ul0DB3HaM2dlkz8fDxTwVRzXmyFmpd+ev+HSdXFv9P1kHz+96TGto2hNbdn0ff83oFUNFg==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"List Prompts","permalink":"/docs/1backend/list-prompts"},"next":{"title":"Remove Prompt","permalink":"/docs/1backend/remove-prompt"}}');var o=t(74848),n=t(28453),i=t(53746),p=t.n(i),a=t(56518),d=t.n(a),c=t(99972),b=t.n(c),m=t(25342),l=t.n(m),h=(t(44215),t(82223),t(24861));const u={id:"subscribe-to-prompt-responses",title:"Subscribe to Prompt Responses by Thread",description:"Subscribe to prompt responses by thread via Server-Sent Events (SSE).",sidebar_label:"Subscribe to Prompt Responses by Thread",hide_title:!0,hide_table_of_contents:!0,api:"eJzdVU1v4zYQ/SvEnHYXtGTXBQro1GwbLIxu0SByDkViYGlqLLGRSC5J2XUF/fdiREt2vEl76qUXm5Lm483Me8MOCvTSKRuU0ZBB3m7pcYssGGadaWxgDr012qNn2yMLlUNRsL0SLEe3RzfLUQd2u0cdPHuX57fvkyf9u2mZFJr5y3DR1bMt7oxDFio8MuGQSYciYJE86XWFzAeHosGCDkqXfjAxOxaOFtmXfPj6U9Xq5y+ceRzCsLsB6Ppo0TPUhTVKB7YzjjWUqMAgVO0T4GAsOkGlrgrIYEK3NjHC/VgpcLDCiQYDOg/ZY3fVpnXswupn4KDohRWhAg5aNAgZxEpXBXBw+LVVDgvIgmuRg5cVNgKyDqggAjGUCX2/IeMxf9bBd/M5/V3NZ6hf6XKaCnCQRgfUgaw/pB/o7+0sfc/h+9ci/6q8p7gjdjY14F8yWEddDSqiRueMey0xH9+Y7R8owwRl8S2UBy3aUBmn/sLiv0vec/AoW6fCcZjwRxQO3U0bKsgeNzSOIEoaPkRysHwvYcOhwVAZok+JYeAJOUAaxTLze3k6+rQbe9mn02TTiXRA+UlBkWCtqylMWhsp6sr4kC0Wy+UPQEBGnDnVHUu9RPsNO0kqTyeTJ2A7U9fmgAXpVzBvhUQmdMGCeUbNhIwMZTtnmkFODx4dVcs+m1LpSVHJSPYKRTHQ4kT3m9O0BmXB1Gth1S94HLqv9M4QThqlkMMosRGKKvaiRv8jUa+tRXBGJ9I0F7HvVixvrTWOmh2bVIVgszRdbIV8Rl2QQwo9v+rCzWqmRVB7ZI2SzlCvlUTPbC3CzriGyqmVRNJQ1k35Pt19ZvtlMn+RzWdpejgcklK3iXFlevLzqShtPVsm86QKTU0YArrG/7bLY7YzWH8QZYkuUSYdTFLqkwo1mSw+xkKAA9Ehwp8ny2Q+czIhDnCwxodG6AukL1b1iaL3l6s6Lim46kt31tP/fdufeBjwz5DaWihNjRyG2p1k+whn2ZKWo3CBQ/ZihZ9vhbN6NxxIpRSj67bC44Or+55ef23R0UrZcNgLp8SWhkxXiPJ0LiDbidrjP8zl3f3p1njPLm+aV+sZxaaPRB9Rt/QEHJ7xeHkT9Zuej8IlMPHzjZRow4UjLdj+csl9ul0DB3HaM2dlkz8fDxTwVRzXmyFmpd+ev+HSdXFv9P1kHz+96TGto2hNbdn0ff83oFUNFg==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},R=void 0,k={},S=[];function v(e){const s={code:"code",p:"p",...(0,n.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Subscribe to Prompt Responses by Thread"}),"\n",(0,o.jsx)(p(),{method:"get",path:"/prompt-svc/prompts/{threadId}/responses/subscribe",context:"endpoint"}),"\n",(0,o.jsxs)(s.p,{children:["Subscribe to prompt responses by thread via Server-Sent Events (SSE).\nYou can subscribe to threads before they are created.\nThe streamed strings are of type ",(0,o.jsx)(s.code,{children:"StreamChunk"}),", see the PromptTypes endpoint for more details."]}),"\n",(0,o.jsx)(h.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,o.jsx)(d(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,o.jsx)(b(),{title:"Body",body:void 0}),"\n",(0,o.jsx)(l(),{id:void 0,label:void 0,responses:{200:{description:"Streaming response",content:{"*/*":{schema:{type:"string"}}}},400:{description:"Missing threadId parameter",content:{"*/*":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"*/*":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function g(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,o.jsx)(s,{...e,children:(0,o.jsx)(v,{...e})}):v(e)}}}]);