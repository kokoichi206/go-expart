//
//  MobileEbitenViewControllerWithErrorHandling.m
//  ios
//
//  Created by Takahiro Tominaga on 2024/03/31.
//

#import "MobileEbitenViewControllerWithErrorHandling.h"

#import <Foundation/Foundation.h>

@implementation MobileEbitenViewControllerWithErrorHandling {
}

- (void)onErrorOnGameUpdate:(NSError*)err {
    // You can define your own error handling e.g., using Crashlytics.
    NSLog(@"Inovation Error!: %@", err);
}

@end

