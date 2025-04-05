"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6918],{50543:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>k,contentTitle:()=>f,default:()=>g,frontMatter:()=>b,metadata:()=>r,toc:()=>w});const r=JSON.parse('{"id":"1backend/stop-container","title":"Stop a Container","description":"Stops a Docker container with the specified parameters.","source":"@site/docs/1backend/stop-container.api.mdx","sourceDirName":"1backend","slug":"/1backend/stop-container","permalink":"/docs/1backend/stop-container","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"stop-container","title":"Stop a Container","description":"Stops a Docker container with the specified parameters.","sidebar_label":"Stop a Container","hide_title":true,"hide_table_of_contents":true,"api":"eJzFVU1z2zYQ/SuYPVOkVDl1h6fababjJlN7rPgkayYrcEUhJgEYACUrGv73zoLUp91MckkvFAntx3sLvIctFOSlUzYooyGHSTDWCxR/GvlETkijAypNTqxVWIqwJOEtSbVQVAiLDmsK5Hz6qB/1PT03ypGPUZ/3mQO/kvn+K/fB2M/CkquV98roFBIwlhwygJsCcuCIP3bxkICj54Z8uDbFBvItcCnSgV/R2krJmJl98Yx/C14uqUZ+s47rBkWev1TBT3rB2lYEOVyML3+bX/5Kw3dzhATCxlLs7ZQuoU1AY02nGb6xgwOrl9dJ7X7FzL+QDNDy0uv5ij07cd9x61kqRwXkwTXU8oK3RvsO/S/DIf+c1rr9AMn3j+MVtDaBi7fK3ugVVqoQf09u//mRBqfzJueMO+r7rSFFJKPXSB40NmFpnPpKxU9D8u7tmQRyGisxIbciJ97Hmj8HUpuAJ9k4FTaQT7dwTejIXTVhCfl01s4SCFh6yKdwOFeTlYRZAjWFpWFN2YbPmEXOgexEm4evjJUH3I05+tiscRVnZJWRWC2ND/loNB5fArfdoZowy47YMbbzGX7aWBKPfcgjiIWpKrOmQsw3AoW3KEmgLkQwT6QFyk4PYuFMHS3lwXfExEdTKi1IF9YoHdhAFNdfEhbRMDrpwlV/duKWHNSKVn2gTZw1b9P9wV3e76TeucW5RfRlz22g5fYLs7MmlPE0UI2KJ+exIv+7V7psKgzO6FSa+gjj3Y2YNNYax/vTDXsZgs2zbDRH+US64IQMXhnJ1c1AY1ArErWSzvCeKUle2ArDwriax1IpSdpHQrt+f919FKtxOjzp5vMsW6/Xaamb1Lgy6/N8hqWtBuN0mC5DXTGGQK72t4tJ1+0A1q+xLMmlymQxJON5qxB9c3TdEYEE+Fh18IfpOB0OnEzHF1zXGh9q1EdIo1GiOL4ITgZwdBH8P5dWf5wCvYTMVqg084gz3fZCm8JJvd4udmyi2GYJsKg4drudo6cHV7UtLz835FjvswRW6BTOeZbTWZvszjmr84k2kEfZkw4DFhgPGaumO+hnhtQmu4wrKcmGb8Yeu8fdwydIYN5fwLUpOMXhmq8tXEMO8Q7n5GgCcW0LFeqywZJju5IsOeyt4SBGRpTsXpjU7i+9OQJ4LuaOBz+Z1Zsp220n9bbdx3d//WfG3kG6aN7SWdu2/wLKOyf/","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Check If a Container Is Running","permalink":"/docs/1backend/container-is-running"},"next":{"title":"Get Container Summary","permalink":"/docs/1backend/container-summary"}}');var i=n(74848),s=n(28453),a=n(53746),o=n.n(a),c=n(56518),p=n.n(c),d=n(99972),l=n.n(d),u=n(25342),m=n.n(u),h=(n(44215),n(82223),n(24861));const b={id:"stop-container",title:"Stop a Container",description:"Stops a Docker container with the specified parameters.",sidebar_label:"Stop a Container",hide_title:!0,hide_table_of_contents:!0,api:"eJzFVU1z2zYQ/SuYPVOkVDl1h6fababjJlN7rPgkayYrcEUhJgEYACUrGv73zoLUp91MckkvFAntx3sLvIctFOSlUzYooyGHSTDWCxR/GvlETkijAypNTqxVWIqwJOEtSbVQVAiLDmsK5Hz6qB/1PT03ypGPUZ/3mQO/kvn+K/fB2M/CkquV98roFBIwlhwygJsCcuCIP3bxkICj54Z8uDbFBvItcCnSgV/R2krJmJl98Yx/C14uqUZ+s47rBkWev1TBT3rB2lYEOVyML3+bX/5Kw3dzhATCxlLs7ZQuoU1AY02nGb6xgwOrl9dJ7X7FzL+QDNDy0uv5ij07cd9x61kqRwXkwTXU8oK3RvsO/S/DIf+c1rr9AMn3j+MVtDaBi7fK3ugVVqoQf09u//mRBqfzJueMO+r7rSFFJKPXSB40NmFpnPpKxU9D8u7tmQRyGisxIbciJ97Hmj8HUpuAJ9k4FTaQT7dwTejIXTVhCfl01s4SCFh6yKdwOFeTlYRZAjWFpWFN2YbPmEXOgexEm4evjJUH3I05+tiscRVnZJWRWC2ND/loNB5fArfdoZowy47YMbbzGX7aWBKPfcgjiIWpKrOmQsw3AoW3KEmgLkQwT6QFyk4PYuFMHS3lwXfExEdTKi1IF9YoHdhAFNdfEhbRMDrpwlV/duKWHNSKVn2gTZw1b9P9wV3e76TeucW5RfRlz22g5fYLs7MmlPE0UI2KJ+exIv+7V7psKgzO6FSa+gjj3Y2YNNYax/vTDXsZgs2zbDRH+US64IQMXhnJ1c1AY1ArErWSzvCeKUle2ArDwriax1IpSdpHQrt+f919FKtxOjzp5vMsW6/Xaamb1Lgy6/N8hqWtBuN0mC5DXTGGQK72t4tJ1+0A1q+xLMmlymQxJON5qxB9c3TdEYEE+Fh18IfpOB0OnEzHF1zXGh9q1EdIo1GiOL4ITgZwdBH8P5dWf5wCvYTMVqg084gz3fZCm8JJvd4udmyi2GYJsKg4drudo6cHV7UtLz835FjvswRW6BTOeZbTWZvszjmr84k2kEfZkw4DFhgPGaumO+hnhtQmu4wrKcmGb8Yeu8fdwydIYN5fwLUpOMXhmq8tXEMO8Q7n5GgCcW0LFeqywZJju5IsOeyt4SBGRpTsXpjU7i+9OQJ4LuaOBz+Z1Zsp220n9bbdx3d//WfG3kG6aN7SWdu2/wLKOyf/",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,k={},w=[];function x(e){const t={code:"code",p:"p",...(0,s.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Stop a Container"}),"\n",(0,i.jsx)(o(),{method:"put",path:"/container-svc/container/stop",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Stops a Docker container with the specified parameters."}),"\n",(0,i.jsxs)(t.p,{children:["Requires the ",(0,i.jsx)(t.code,{children:"container-svc:container:stop"})," permission."]}),"\n",(0,i.jsx)(h.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(p(),{parameters:void 0}),"\n",(0,i.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{id:{example:"4378b76e05ba",type:"string"},name:{example:"sup-container-x",type:"string"}},type:"object"}}},description:"Stop Container Request",required:!0}}),"\n",(0,i.jsx)(m(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function g(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(x,{...e})}):x(e)}}}]);