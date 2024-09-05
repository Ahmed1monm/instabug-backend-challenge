class ApplicationsController < ApplicationController
  def index
    @applications = Application.all
    render json: @applications, status: :ok
  end

  def show
    @application = Application.find_by(token: params[:token])
    render json: @application, status: :ok
  end

  def new
    @application = Application.new
  end

  def create
    @application = Application.new(application_params)

    if @application.save
      render json: @application, status: :created
    else
      render json: @application.errors, status: :unprocessable_entity
    end
  end

  def update
    @application = Application.find_by(token: params[:token])
    
    if @application.update(application_params)
      render json: @application, except: :id, status: :ok
    else
      render json: @application.errors, status: :unprocessable_entity
    end
  end

  private

  def application_params
    params.require(:application).permit(:name)
  end
end
