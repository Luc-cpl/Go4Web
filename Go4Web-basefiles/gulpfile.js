var gulp = require('gulp');

var del = require('del');
var babel = require('gulp-babel');
var uglify = require('gulp-uglify');
var sass = require('gulp-sass');
var cleanCSS = require('gulp-clean-css');
var htmlmin = require('gulp-htmlmin');
var webpack = require('webpack-stream');
var rename = require('gulp-rename');

var babelPreset = 'latest';

gulp.task('default', ['js', 'css', 'html', 'js-public', 'public-css'], function(){
    gulp.watch('views/files/**/*.js',['js']);
    gulp.watch('views/files/**/*.scss',['css']);
    gulp.watch('views/files/**/*.sass',['css']);
    gulp.watch('views/files/**/*.sass',['css']);
    gulp.watch('views/files/**/*.html',['html']);
    gulp.watch('public/**/*.css',['public-css']);
    gulp.watch('public/**/*.js',['public-css']);
});

gulp.task('js', function () {
    del('views/files-min/**/*.js', {force:true});
    return gulp.src(['views/files/**/*.js'])
        .pipe(babel({
            presets: [babelPreset]
        }))
        .pipe(uglify())
        .pipe(gulp.dest('views/files-min'));
});

gulp.task('js-public', function () {
    return gulp.src(['public/**/*.js'])
        .pipe(babel({
            presets: [babelPreset]
        }))
        .pipe(uglify())
        .pipe(rename({ suffix: '.min' }))
        .pipe(gulp.dest('public/min/'));
});

gulp.task('css', function(){
    del('views/files-min/**/*.css', {force:true});
    gulp.src(['views/files/**/*.css'])
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('views/files-min'));
    gulp.src(['views/files/**/*.sass'])
        .pipe(sass().on('error', sass.logError))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('views/files-min'));
    return gulp.src(['views/files/**/*.scss'])
        .pipe(sass().on('error', sass.logError))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('views/files-min'));
});

gulp.task('public-css', function(){
    gulp.src(['public/**/*.css'])
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(rename({ suffix: '.min' }))
        .pipe(gulp.dest('public/'));
    gulp.src(['public/**/*.sass'])
        .pipe(sass().on('error', sass.logError))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(rename({ suffix: '.min' }))
        .pipe(gulp.dest('public/'));
    return gulp.src(['public/**/*.scss'])
        .pipe(sass().on('error', sass.logError))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(rename({ suffix: '.min' }))
        .pipe(gulp.dest('public/'));
});

gulp.task('html', function(){
    del('views/files-min/**/*.html', {force:true});
    return gulp.src(['views/files/**/*.html'])
        .pipe(htmlmin({collapseWhitespace: true}))
        .pipe(gulp.dest('views/files-min'));
});