"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5928],{45525:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>o,contentTitle:()=>a,default:()=>h,frontMatter:()=>c,metadata:()=>r,toc:()=>l});const r=JSON.parse('{"id":"writing-custom-services/your-first-service","title":"Your first service","description":"While 1Backend itself is written in Go, services that run on 1Backend can be written in any language.","source":"@site/docs/writing-custom-services/your-first-service.md","sourceDirName":"writing-custom-services","slug":"/writing-custom-services/your-first-service","permalink":"/docs/writing-custom-services/your-first-service","draft":false,"unlisted":false,"editUrl":"https://github.com/1backend/1backend/tree/main/docs-source/docs/writing-custom-services/your-first-service.md","tags":[{"inline":true,"label":"test","permalink":"/docs/tags/test"}],"version":"current","sidebarPosition":3,"frontMatter":{"sidebar_position":3,"tags":["test"]},"sidebar":"tutorialSidebar","previous":{"title":"Writing custom services","permalink":"/docs/category/writing-custom-services"},"next":{"title":"Command line","permalink":"/docs/category/command-line"}}');var s=n(74848),i=n(28453);const c={sidebar_position:3,tags:["test"]},a="Your first service",o={},l=[{value:"A Go example",id:"a-go-example",level:2},{value:"Things to understand",id:"things-to-understand",level:2},{value:"Instance registration",id:"instance-registration",level:3}];function d(e){const t={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",header:"header",li:"li",p:"p",pre:"pre",ul:"ul",...(0,i.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(t.header,{children:(0,s.jsx)(t.h1,{id:"your-first-service",children:"Your first service"})}),"\n",(0,s.jsx)(t.p,{children:"While 1Backend itself is written in Go, services that run on 1Backend can be written in any language.\nA service only needs a few things to fully function:"}),"\n",(0,s.jsxs)(t.ul,{children:["\n",(0,s.jsxs)(t.li,{children:["Register a user account, just like a human user. For details, see the ",(0,s.jsx)(t.a,{href:"/docs/built-in-services/user-svc",children:"User Svc"}),"."]}),"\n",(0,s.jsx)(t.li,{children:"Register its instance in the registry so 1Backend knows where to route requests."}),"\n"]}),"\n",(0,s.jsx)(t.h2,{id:"a-go-example",children:"A Go example"}),"\n",(0,s.jsx)(t.p,{children:"The following Go service demonstrates these steps:"}),"\n",(0,s.jsxs)(t.ul,{children:["\n",(0,s.jsxs)(t.li,{children:["Registers itself as a user with the slug ",(0,s.jsx)(t.code,{children:"basic-svc"})]}),"\n",(0,s.jsxs)(t.li,{children:["Registers or updates its URL (",(0,s.jsx)(t.code,{children:"http://127.0.0.1:9111"}),") in the ",(0,s.jsx)(t.a,{href:"/docs/built-in-services/registry-svc",children:"Registry"}),"."]}),"\n"]}),"\n",(0,s.jsx)(t.p,{children:'You may notice that the following code uses a "Go SDK," but it\'s simply a set of convenience functions built on top of the 1Backend API.\n1Backend is language-agnostic and can be used with any language, even if no SDK is available in the repository.'}),"\n",(0,s.jsx)(t.pre,{children:(0,s.jsx)(t.code,{className:"language-go",children:'// \x3c!-- INCLUDE: ../../../examples/go/services/basic/internal/basic_service.go --\x3e\npackage basicservice\n\nimport (\n\t"context"\n\t"net/http"\n\t"os"\n\n\topenapi "github.com/1backend/1backend/clients/go"\n\tbasic "github.com/1backend/1backend/examples/go/services/basic/internal/types"\n\tsdk "github.com/1backend/1backend/sdk/go"\n\t"github.com/1backend/1backend/sdk/go/auth"\n\t"github.com/1backend/1backend/sdk/go/boot"\n\t"github.com/1backend/1backend/sdk/go/client"\n\t"github.com/1backend/1backend/sdk/go/datastore"\n\t"github.com/1backend/1backend/sdk/go/infra"\n\t"github.com/1backend/1backend/sdk/go/middlewares"\n\t"github.com/gorilla/mux"\n\t"github.com/pkg/errors"\n)\n\nconst RolePetManager = "basic-svc:pet:manager"\n\ntype BasicService struct {\n\tOptions *Options\n\n\ttoken            string\n\tuserSvcPublicKey string\n\n\tdataStoreFactory infra.DataStoreFactory\n\n\tpetsStore       datastore.DataStore\n\tcredentialStore datastore.DataStore\n\n\tRouter *mux.Router\n}\n\ntype Options struct {\n\tTest      bool\n\tServerUrl string\n\tSelfUrl   string\n}\n\nfunc NewService(options *Options) (*BasicService, error) {\n\tif options.ServerUrl == "" {\n\t\toptions.ServerUrl = os.Getenv("OB_SERVER_URL")\n\t}\n\tif options.ServerUrl == "" {\n\t\toptions.ServerUrl = "http://127.0.0.1:58231"\n\t}\n\tif options.SelfUrl == "" {\n\t\toptions.SelfUrl = os.Getenv("OB_SELF_URL")\n\t}\n\n\tdconf := infra.DataStoreConfig{}\n\tif options.Test {\n\t\tdconf.TablePrefix = sdk.Id("t")\n\t}\n\n\tservice := &BasicService{\n\t\tOptions: options,\n\t}\n\n\tdsf, err := infra.NewDataStoreFactory(dconf)\n\tif err != nil {\n\t\treturn nil, errors.Wrap(err, "cannot create datastore factory")\n\t}\n\tservice.dataStoreFactory = dsf\n\n\tpetStore, err := dsf.Create("basicSvcPets", &basic.Pet{})\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tservice.petsStore = petStore\n\n\tservice.registerAccount()\n\tservice.registerRoutes()\n\n\treturn service, nil\n}\n\nfunc (service *BasicService) Start() error {\n\tclient := client.NewApiClientFactory(service.Options.ServerUrl).\n\t\tClient(client.WithToken(service.token))\n\n\t_, _, err := client.RegistrySvcAPI.\n\t\tRegisterInstance(context.Background()).\n\t\tBody(openapi.RegistrySvcRegisterInstanceRequest{\n\t\t\tUrl: service.Options.SelfUrl,\n\t\t}).Execute()\n\tif err != nil {\n\t\treturn errors.Wrap(err, "cannot register instance")\n\t}\n\n\treturn nil\n}\n\nfunc (service *BasicService) registerAccount() error {\n\tcredentialStore, err := service.dataStoreFactory.Create("basicSvcCredentials", &auth.Credential{})\n\tif err != nil {\n\t\treturn errors.Wrap(err, "cannot create credential store")\n\t}\n\tservice.credentialStore = credentialStore\n\n\tobClient := client.NewApiClientFactory(service.Options.ServerUrl).Client()\n\ttoken, err := boot.RegisterServiceAccount(\n\t\tobClient.UserSvcAPI,\n\t\t"basic-svc",\n\t\t"Basic Svc",\n\t\tservice.credentialStore,\n\t)\n\tif err != nil {\n\t\treturn errors.Wrap(err, "cannot register service")\n\t}\n\tservice.token = token.Token\n\n\tobClient = client.NewApiClientFactory(service.Options.ServerUrl).\n\t\tClient(client.WithToken(service.token))\n\n\t_, _, err = obClient.RegistrySvcAPI.\n\t\tRegisterInstance(context.Background()).\n\t\tBody(openapi.RegistrySvcRegisterInstanceRequest{\n\t\t\tUrl: service.Options.SelfUrl,\n\t\t}).Execute()\n\tif err != nil {\n\t\treturn errors.Wrap(err, "cannot register instance")\n\t}\n\n\tpk, _, err := obClient.\n\t\tUserSvcAPI.GetPublicKey(context.Background()).\n\t\tExecute()\n\tif err != nil {\n\t\treturn err\n\t}\n\tservice.userSvcPublicKey = pk.PublicKey\n\n\treturn nil\n}\n\nfunc (service *BasicService) registerRoutes() {\n\tmws := []middlewares.Middleware{\n\t\tmiddlewares.ThrottledLogger,\n\t\tmiddlewares.Recover,\n\t\tmiddlewares.CORS,\n\t\tmiddlewares.GzipDecodeMiddleware,\n\t}\n\tappl := applicator(mws)\n\n\tservice.Router = mux.NewRouter()\n\n\tservice.Router.HandleFunc("/basic-svc/pet", appl(func(w http.ResponseWriter, r *http.Request) {\n\t\tservice.SavePet(w, r)\n\t})).\n\t\tMethods("OPTIONS", "PUT")\n\n\tservice.Router.HandleFunc("/basic-svc/pets", appl(func(w http.ResponseWriter, r *http.Request) {\n\t\tservice.ListPets(w, r)\n\t})).\n\t\tMethods("OPTIONS", "POST")\n}\n\nfunc applicator(\n\tmws []middlewares.Middleware,\n) func(http.HandlerFunc) http.HandlerFunc {\n\treturn func(h http.HandlerFunc) http.HandlerFunc {\n\t\tfor _, middleware := range mws {\n\t\t\th = middleware(h)\n\t\t}\n\n\t\treturn h\n\t}\n}\n// \x3c!-- /INCLUDE --\x3e\n'})}),"\n",(0,s.jsx)(t.p,{children:"Just make sure you run it with the appropriate envars:"}),"\n",(0,s.jsx)(t.pre,{children:(0,s.jsx)(t.code,{className:"language-sh",children:"OB_SERVER_URL=http://127.0.0.1:58231 OB_SELF_URL=http://127.0.0.1:9111 go run main.go\n"})}),"\n",(0,s.jsx)(t.p,{children:"Once it's running you will be able to call the 1Backend server proxy and that will proxy to your basic service:"}),"\n",(0,s.jsx)(t.pre,{children:(0,s.jsx)(t.code,{className:"language-sh",children:'# 127.0.0.1:58231 here is the address of the 1Backend server\n$ curl 127.0.0.1:58231/basic-svc/hello\n{"hello": "world"}\n'})}),"\n",(0,s.jsx)(t.p,{children:"This is so you don't have to expose your basic service to the outside world, only your 1Backend server."}),"\n",(0,s.jsx)(t.p,{children:"Let's recap how the proxying works:"}),"\n",(0,s.jsxs)(t.ul,{children:["\n",(0,s.jsxs)(t.li,{children:["Service registers an account, acquires the ",(0,s.jsx)(t.code,{children:"basic-svc"})," slug."]}),"\n",(0,s.jsxs)(t.li,{children:["Service calls the 1Backend ",(0,s.jsx)(t.a,{href:"/docs/built-in-services/registry-svc",children:"Registry Svc"})," to tell the system an instance of the Basic service is available under the URL ",(0,s.jsx)(t.code,{children:"http://127.0.0.1:9111"})]}),"\n",(0,s.jsxs)(t.li,{children:["When you curl the 1Backend server with a path like ",(0,s.jsx)(t.code,{children:"127.0.0.1:58231/basic-svc/hello"}),", the first section of the path will be a user account slug. The daemon checks what instances are owned by that slug and routes the request to one of the instances."]}),"\n"]}),"\n",(0,s.jsx)(t.pre,{children:(0,s.jsx)(t.code,{className:"language-sh",children:"$ oo instance ls\nID                URL                     STATUS    OWNER SLUG       LAST HEARTBEAT\ninst_eHFTNvAlk9   http://127.0.0.1:9111   Healthy   basic-svc     10s ago\n"})}),"\n",(0,s.jsx)(t.h2,{id:"things-to-understand",children:"Things to understand"}),"\n",(0,s.jsx)(t.h3,{id:"instance-registration",children:"Instance registration"}),"\n",(0,s.jsxs)(t.p,{children:["Like most other things on the platform, service instances become owned by a user account slug. When the basic service calls ",(0,s.jsx)(t.a,{href:"/docs/1backend/register-instance",children:"RegisterInstance"}),", the host will be associated with the ",(0,s.jsx)(t.code,{children:"basic-svc"})," slug."]}),"\n",(0,s.jsx)(t.p,{children:"Updates to this host won't be possible unless the caller is the basic service (or the caller is an admin). The service becomes the owner of that URL essentially."}),"\n",(0,s.jsx)(t.p,{children:"This is the same ownership model like in other parts of the 1Backend system."})]})}function h(e={}){const{wrapper:t}={...(0,i.R)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(d,{...e})}):d(e)}}}]);