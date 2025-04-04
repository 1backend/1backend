"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1663],{15974:(e,t,o)=>{o.r(t),o.d(t,{assets:()=>P,contentTitle:()=>f,default:()=>b,frontMatter:()=>k,metadata:()=>r,toc:()=>m});const r=JSON.parse('{"id":"1backend/checkout-repo","title":"Checkout a git repository","description":"Checkout a git repository over https or ssh at a specific version into a temporary directory.","source":"@site/docs/1backend/checkout-repo.api.mdx","sourceDirName":"1backend","slug":"/1backend/checkout-repo","permalink":"/docs/1backend/checkout-repo","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"checkout-repo","title":"Checkout a git repository","description":"Checkout a git repository over https or ssh at a specific version into a temporary directory.","sidebar_label":"Checkout a git repository","hide_title":true,"hide_table_of_contents":true,"api":"eJzVVt9v2zYQ/lcIPnWAIiXNCgx6WtJ1S7ZgMeLkKTGKM3WW2FAkS1J2VUP/+3CUbCuxEwR9KLAXQz7eb373Hde8QC+ctEEazXP+sULxaJrAgJUyMIfWeBmMa5lZomNVCNYz45j3FQPS8haFXEjBlui8NJpJHQwDFrC2xoFrWSEdCnKRPugJuoVxtSfDCpQyKyaU0chWMlSsllrWoFglfQy5MI4twAd0TAx5pTzhxqIDyvey4DnfnNygNTzhDr826MO5KVqer7kwOqAO9AnWKimiYfbFU7Vr7kWFNdCXdeQ2SPTxH3i/Mq6g76cNmgwn1IRgHlHHLC9ubydTBk2oeMJDa5Hn3Acndcm7hHtffX7Edt/ZdHrBrJNLCMgesWXvTDwBFZ3SqTBaoyDhL694/mxXr6W6cTYOJRcMtXCtDVjs4h4MEsvcd397oPr9Cg56bJza9/dno9QYcHc3V+wdpmWa9LDLs6yUoWrmqTB11nh0GWkfDuDRaahxP8rdcDJKfOgO2bwx/wHr+97PHWhRJSxAmZBfYepaBja9ONt3020lZv4FReAdiV4YR0I3u+mxPaBcOix4HlyDHQm8Ndr36H1/fHwAa40Q6P2iUartxwkLRq5DhaO28+RHZ6aQbj/qH5vhZ6sKHT4Lxlbgx7m8rUldwn89VOGlXoKSBft7ev3vj5eBzplYyFszOTkAMk3DYJz8jsVPy+TD4Z4EwrtiU3RE4J+iz5+TEvETisbJ0PL8fs3PERy6MyLJ/H7WzRIeoPQ8v+dT0ziBbLoUfJbwGkNliNqtiWC3QBY881HryC9FHPxsw/yc4lB1PoaJ5MKzTBkBqjI+5B9+e396wingJp8p1deXNM5qj+Jai+xhUHngbGFoZWHB5m1cfSCQgS6GRQCiH0q2cKaOSCeyoarYlSmlZqgLa6SOS0yS/wqhQLqOnqv42YCaeBm7YQAr/8E2dpku6Ga34j59g9oqfLqyNrcyWjx7on5j7MQDx+8EfRd3f7eEupNtWXDEjFIvzGbvgoj4whpk9AUK/e9e6rJREJzRxOOj2ieXbNpYa1zYhifez7PsZA7iEXURiZ/vkeTZ5ZGGIJfIaimcISxIgZ5ZBYFeG9RuJQVqHxu1iffX5IotT9PjJ9Foy6xWq7TUTWpcmQ12PoPSqqPT9DitQq3iVkRX++vFtI+2S9avoCzRpdJkUSWj7spAl8RPzvtCnjTvOD1Nj4+cSAmjSQR9DXqU6YtvMv6sE6Pnzv/iITfAO+C3kFkFUo/eBv3Q3/Pd0Me1F99428GfJZwGnPTW6zl4vHOq60j8tUFHrDNL+BKchDn1/37WJZuZI6boR+Nj37MjGna6GFBNP3TPaLFLNhZnQqANr+qOaWxyPb3lCZ8PT9LaFGTjYEUVwYrnPD5qyToyUpStuQJdNlCSbu+T5h8GntoxA6WUbD6oqs2RbkcZPmeWvhD6pbIOmqzXPe903Va/P3rRYktnvTbd56zruv8AkFNPcg==","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Save Secrets","permalink":"/docs/1backend/save-secrets"},"next":{"title":"Reset Password","permalink":"/docs/1backend/reset-password"}}');var i=o(74848),s=o(28453),a=o(53746),n=o.n(a),c=o(56518),p=o.n(c),d=o(99972),l=o.n(d),h=o(25342),y=o.n(h),u=(o(44215),o(82223),o(24861));const k={id:"checkout-repo",title:"Checkout a git repository",description:"Checkout a git repository over https or ssh at a specific version into a temporary directory.",sidebar_label:"Checkout a git repository",hide_title:!0,hide_table_of_contents:!0,api:"eJzVVt9v2zYQ/lcIPnWAIiXNCgx6WtJ1S7ZgMeLkKTGKM3WW2FAkS1J2VUP/+3CUbCuxEwR9KLAXQz7eb373Hde8QC+ctEEazXP+sULxaJrAgJUyMIfWeBmMa5lZomNVCNYz45j3FQPS8haFXEjBlui8NJpJHQwDFrC2xoFrWSEdCnKRPugJuoVxtSfDCpQyKyaU0chWMlSsllrWoFglfQy5MI4twAd0TAx5pTzhxqIDyvey4DnfnNygNTzhDr826MO5KVqer7kwOqAO9AnWKimiYfbFU7Vr7kWFNdCXdeQ2SPTxH3i/Mq6g76cNmgwn1IRgHlHHLC9ubydTBk2oeMJDa5Hn3Acndcm7hHtffX7Edt/ZdHrBrJNLCMgesWXvTDwBFZ3SqTBaoyDhL694/mxXr6W6cTYOJRcMtXCtDVjs4h4MEsvcd397oPr9Cg56bJza9/dno9QYcHc3V+wdpmWa9LDLs6yUoWrmqTB11nh0GWkfDuDRaahxP8rdcDJKfOgO2bwx/wHr+97PHWhRJSxAmZBfYepaBja9ONt3020lZv4FReAdiV4YR0I3u+mxPaBcOix4HlyDHQm8Ndr36H1/fHwAa40Q6P2iUartxwkLRq5DhaO28+RHZ6aQbj/qH5vhZ6sKHT4Lxlbgx7m8rUldwn89VOGlXoKSBft7ev3vj5eBzplYyFszOTkAMk3DYJz8jsVPy+TD4Z4EwrtiU3RE4J+iz5+TEvETisbJ0PL8fs3PERy6MyLJ/H7WzRIeoPQ8v+dT0ziBbLoUfJbwGkNliNqtiWC3QBY881HryC9FHPxsw/yc4lB1PoaJ5MKzTBkBqjI+5B9+e396wingJp8p1deXNM5qj+Jai+xhUHngbGFoZWHB5m1cfSCQgS6GRQCiH0q2cKaOSCeyoarYlSmlZqgLa6SOS0yS/wqhQLqOnqv42YCaeBm7YQAr/8E2dpku6Ga34j59g9oqfLqyNrcyWjx7on5j7MQDx+8EfRd3f7eEupNtWXDEjFIvzGbvgoj4whpk9AUK/e9e6rJREJzRxOOj2ieXbNpYa1zYhifez7PsZA7iEXURiZ/vkeTZ5ZGGIJfIaimcISxIgZ5ZBYFeG9RuJQVqHxu1iffX5IotT9PjJ9Foy6xWq7TUTWpcmQ12PoPSqqPT9DitQq3iVkRX++vFtI+2S9avoCzRpdJkUSWj7spAl8RPzvtCnjTvOD1Nj4+cSAmjSQR9DXqU6YtvMv6sE6Pnzv/iITfAO+C3kFkFUo/eBv3Q3/Pd0Me1F99428GfJZwGnPTW6zl4vHOq60j8tUFHrDNL+BKchDn1/37WJZuZI6boR+Nj37MjGna6GFBNP3TPaLFLNhZnQqANr+qOaWxyPb3lCZ8PT9LaFGTjYEUVwYrnPD5qyToyUpStuQJdNlCSbu+T5h8GntoxA6WUbD6oqs2RbkcZPmeWvhD6pbIOmqzXPe903Va/P3rRYktnvTbd56zruv8AkFNPcg==",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,P={},m=[];function v(e){const t={p:"p",...(0,s.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"Checkout a git repository"}),"\n",(0,i.jsx)(n(),{method:"post",path:"/source-svc/repo/checkout",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Checkout a git repository over https or ssh at a specific version into a temporary directory.\nPerforms a shallow clone with minimal history for faster checkout."}),"\n",(0,i.jsx)(u.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(p(),{parameters:void 0}),"\n",(0,i.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{password:{description:"Password or token for HTTPS auth",type:"string"},ssh_key:{description:"SSH private key (optional for SSH connection)",type:"string"},ssh_key_pwd:{description:"Password for SSH private key if encrypted (optional)",type:"string"},token:{description:"Token for HTTPS auth (optional for SSH)",type:"string"},url:{description:"Full repository URL (e.g., https://github.com/user/repo)",type:"string"},username:{description:"Username for HTTPS or SSH user (optional for SSH)",type:"string"},version:{description:"Branch, tag, or commit SHA",type:"string"}},type:"object"}}},description:"Checkout Repo Request",required:!0}}),"\n",(0,i.jsx)(y(),{id:void 0,label:void 0,responses:{200:{description:"Successfully checked out the repository",content:{"application/json":{schema:{properties:{dir:{description:"Directory where the repository was checked out",type:"string"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function b(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(v,{...e})}):v(e)}}}]);