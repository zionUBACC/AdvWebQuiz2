module Main exposing (main)
import Html exposing (..)
import Html.Attributes exposing (class, src)
import Html.Events exposing (onClick)
import Browser


baseUrl : String
baseUrl = 
    "images/"
    
    
viewDetailedPhoto : {liked : Bool } -> Html Msg
viewDetailedPhoto model =
    let
        buttonType =
            if model.liked then
                "star"
            else
                "star_border"
        msg = 
            if model.liked then
                Unlike
            else
                Like
    in

        div [class "main-div"]
            [ 
                div [ class "like-button" ]
                    [ span
                        [ class "material-icons md-48"
                        , onClick msg 
                        ]
                        [  text buttonType ]
                    ]
            ]
        
            
            
            
            
--MODEL
initialModel : {liked : Bool }
initialModel =
    { 
        liked = False
    }



type Msg = Like | Unlike




update : Msg -> {liked : Bool} -> {liked : Bool}
update msg model =
    case msg of 
        Like ->
            { model | liked = True }
        Unlike ->
            { model | liked = False }



view : {liked : Bool} -> Html Msg
view model = 
    div []
    [ 
        div [ class "header" ] 
        [ h1[] [ text "Quiz" ] ]
    ,div [ class "content-flow" ]
        [ viewDetailedPhoto model ]
    ]



main : Program () {liked : Bool} Msg
main =
    Browser.sandbox
    {
        init = initialModel
        ,view = view
        ,update = update
    }
   

