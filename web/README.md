# React-Develop-Template

## 特点

- 1、基于 TypeScript
- 2、基于最新的 React 18
- 3、基于最流行的设计风格 Ant Design v5.x
- 4、基于 React Router v6.x 做路由管理，支持懒加载
- 5、基于 Vite4 做项目编译打包工具
- 6、基于 Redux、Redux Toolkit 做状态管理
- 7、基于 RTK Query 请求管理
- 8、完善的 **国际化** 配置支持
- 9、完善的 Mock 数据支持

## 开始使用

**Clone**

```shell
git clone git@github.com:lxhanghub/react-develop-template.git
```

**Install**

```shell
cd react-develop-template
npm install
```

**Run**

```shell
vite
```

**Build**

```shell
# 开发环境
npm run build:dev

# 测试环境
npm run build:test

# 生产环境
npm run build:pro
```

## 目录结构

```
.
├── docs                  # 文档内容
├── mock
│    └── api.mock.ts      # 开发环境的 Mock 数据定义
├── public                # 静态资源文件目录 
├── src
│    ├── apis             # API 定义目录
│    ├── assets           # 资源文件
│    ├── components       # 通用组件定义
│    ├── context          # React Context
│    ├── hooks            # React 自定义 Hook
│    ├── layout           # 布局文件以及布局涉及的组件
│    ├── locales          # 国际化语言定义
│    ├── pages            # 页面文件夹
│    ├── routers          # 路由和菜单的定义
│    ├── store            # redux store 定义
│    ├── App.tsx          # React 运行入口文件
│    ├── main.tsx         # 入口文件
│    └── vite-env.d.ts    # Vite 声明文件
├── index.html            # 应用运行入口文件
├── LICENSE               # 授权文件（MIT）
├── package-lock.json     # 依赖包版本锁定文件
├── package.json          # NPM 管理
├── readme.md        
├── tsconfig.json         # TypeScript 配置文件
├── tsconfig.node.json
├── vite.config.ts        # Vite 配置文件
```

## 常见问题
