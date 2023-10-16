FROM ruby:3.2.2-slim

RUN apt-get update -qq && \
    apt-get install -y \
    build-essential \
    libsqlite3-dev \
    libpq-dev \
    git

ENV app /app

RUN mkdir $app

WORKDIR $app

RUN gem install bundler

COPY Gemfile /app/

ENV BUNDLE_PATH /box
