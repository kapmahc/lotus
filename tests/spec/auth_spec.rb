require 'spec_helper'
require 'airborne'

describe :front do
  describe :auth do
    it 'GET /layout' do
      get "#{ENV['HOST']}/layout"
      expect_json_types(title: :string)
    end

    it 'GET /users/sign-in' do
    end

    it 'POST /users/sign-in' do
    end

    it 'GET /users/sign-up' do
    end

    it 'POST /users/sign-up' do
    end
  end
end
