require('ts-node').register({
    lazy: true,
    ignore: [/node_modules\//],
    project: './.gulp/tsconfig.gulp.json'
});

require('./.gulp/gulpfile');