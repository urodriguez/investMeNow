## Usage

``` bash
# serve with hot reload at localhost:8080
npm run serve

# build for production with minification
npm run build

# run linter
npm run lint

# run unit tests
npm run test:unit

# run e2e tests
npm run test:e2e

```

For a detailed explanation on how things work, check out the [Vue CLI Guide](https://cli.vuejs.org/guide/).

## What's included

Within the download you'll find the following directories and files:

```
CoreUI-Vue/
├── public/              # pure static assets (directly copied)
│   └── index.html           # index.html template
├── src/                 # project root
│   ├── assets/                 # module assets (processed by webpack)
│   │   └── scss/               # user styles
│   ├── components/             # ui components
│   ├── containers/             # ui containers
│   ├── router/                 # routing 
│   ├── shared/                 # utils
│   ├── views/                  # ui views
│   ├── _nav.js                 # sidebar nav config
│   ├── App.vue                 # main app component
│   └── main.js                 # app entry file
├── test/
│   └── unit/            # unit tests
│   └── e2e/             # e2e tests
├── .eslintrc.js         # eslint config
├── .gitignore           # defaults for gitignore
├── .postcssrc.js        # postcss config
├── CHANGELOG.md
├── README.md
├── babel.config.js      # babel config
├── jest.config.js       # jest config
├── vue.config.js        # vue-cli config
├── LICENSE
└── package.json         # build scripts and dependencies
```