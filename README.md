# Image2ascii Converter
```
./main <path/to/file>.png
```

## 処理の流れ
1. 画像リサイズ（別ファイル出力）
2. 画像をグレースケールに変更
3. 色の強度値を文字に変換する
4. 出力（output.txt）

## 注意点
- 100*100にリサイズされるので縦横比が違う画像だと引き伸ばされる
- 今のところpngしか対応してねぇわ
- 画像によってはうまくいかない


## 出力サンプル
でかっ

```
                                                                                                    
                                                                                                    
                                                                MM MMM                              
                 MMM                                          M     MM                              
               MMM   M                                      M                                       
               MMM                                         M                                        
                MM                                        M                                         
                        M                                M                                          
                         M                                                                          
                                   MMM         M                                                    
                          M                       MM   M                                            
                           M    M                    MM                                             
                              M                                                                     
                            M                        M                                              
                         MM                                 M                                       
                        M                                     M                                     
                                                               M                                    
                      M                                         M                                   
                     M                                           M                                  
                                      M       MM                  M                                 
                    M                            M                 M                                
                   M             M                                  M                               
                   M           M                    M                                               
                  M          M                       M               M                              
                 M          M                         M               M                             
                 M         M                                                                        
                M                                      M               M                            
                         M                   MMMMM     M               M                            
               M              MMMM          M     M     M               M                           
                        M         M        M            M                                           
              M             M     M                M                     M                          
                       M         MMM            MMMM     M               M                          
             M         M   M     M M      M     MMMMM    M                M                         
             M         M  M      MMM      M     MMM M                     M                         
                          M        M      M         M     M                                         
            M             M        M                M                      M                        
                          M                         M                      M                        
                      M   M        M       M        M      M                                        
           M          M   M       M         M     M        M                M                       
           M          M          M                                                                  
           M          M                                                      M                      
                      M                                                      M                      
                      M                                                       M                     
          M                          M                                        M                     
          M                                                                                         
          M            M                                   M                   M                    
                                                           M                   M                    
         M              M                                                      M                    
         M                                                M                                         
                         M         M                     M                                          
                          M                             M                       M                   
                           M                          MM                        M                   
                                                    MM                                              
                             M                   MM                              M                  
                               M            MM                                   M                  
         M                       MMMMMMMMM                     M                 M                  
         M                                                     M                                    
         M                                                                                          
         M                                                      M                 M                 
         M                                                      M                 M                 
         M          M                                                                               
         M          M                                            M                                  
         M                                                                         M                
         M                                                        M                M                
         M           M                                             M                                
         M           M                                             M                                
         M           M                                             MM               M               
         M         M                                                                                
         M            M                                              M               M              
         M        M   M                                           M   M              M              
                 M    M                                                              M              
           M                                                           M             M              
                       M                                                M            M              
                       M                                         M       M           M              
                                                                 M                                  
                        M                                                   MMMMMMMM                
                        M                                                                           
                                      M                                                             
                                       M                                                            
                                         MMMM                                                       
                                               M                                                    
                                                M                                                   
                         M              M                                                           
                         M              M        M               M                                  
                                                 M                                                  
                                                                                                    
                                                  M                                                 
                                                  M               M                                 
                                                  M               M                                 
                                                                   M                                
                                                   M               M                                
                                       M           M                                                
                                                    M                                               
                         M                          M              M                                
                                                     M             M                                
                                      M                MMMMM     M                                  
                                     M                                                              
                          MM       MM                                                               
                                                                                                    
                                                                                                    
```
