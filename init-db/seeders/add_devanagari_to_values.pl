#!/usr/bin/perl
use strict;
use warnings;

my %a2d = ('0'=>'०','1'=>'१','2'=>'२','3'=>'३','4'=>'४','5'=>'५','6'=>'६','7'=>'७','8'=>'८','9'=>'९');

my $in_translations = 0;

while (<>) {
    my $line = $_;
    
    # Track if we're in getBanglaTranslations function
    if ($line =~ /func.*getBanglaTranslations/) {
        $in_translations = 1;
    } elsif ($line =~ /^func/ && $in_translations) {
        $in_translations = 0;
    }
    
    # Within getBanglaTranslations: convert ASCII to Devanagari in VALUES only
    if ($in_translations && $line =~ /^(\s*)"[^"]*":\s*"/) {
        # Extract value part (everything after ": ")
        if ($line =~ /^(\s*"[^"]*":\s*")(.*?)(".*)$/) {
            my ($prefix, $value, $suffix) = ($1, $2, $3);
            # Convert ASCII to Devanagari in value
            foreach my $a (keys %a2d) {
                $value =~ s/$a/$a2d{$a}/g;
            }
            $line = $prefix . $value . $suffix . "\n";
        }
    }
    
    print $line;
}
